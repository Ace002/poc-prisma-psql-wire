package handler

import (
	"context"
	"fmt"

	"github.com/Ace002/poc-prisma-psql-wire/psql_wire_app/config"
	wire "github.com/jeroenrinzema/psql-wire"
	"github.com/lib/pq/oid"
)

// ColumnInfo holds metadata for each result column
type ColumnInfo struct {
	Name        string
	DataTypeOID uint32 // PostgreSQL data type ID (OID)
}

// Response is the full payload received from the external API after query execution
type Response struct {
	Data             [][]any      `json:"data"`              // Row data
	ColumnDefinition []ColumnInfo `json:"column_definition"` // Column metadata
}

var table = wire.Columns{
	{
		Table: 0,
		Name:  "name",
		Oid:   oid.T__text,
		Width: 256,
	},
	{
		Table: 0,
		Name:  "number",
		Oid:   oid.T__numeric,
		Width: 1,
	},
}

// QueryHandler returns a PostgreSQL wire-compatible handler that executes remote queries
func QueryHandler(cfg *config.Config) wire.ParseFn {
	return func(ctx context.Context, query string) (wire.PreparedStatements, error) {
		// log := logger.Get()

		// Define the statement with parameter support
		stmt := wire.NewStatement(func(ctx context.Context, writer wire.DataWriter, parameters []wire.Parameter) error {
			defer func() {
				if r := recover(); r != nil {
					// log.Error("🔥 Panic in handler", "error", r, "stack", string(debug.Stack()))
				}
			}()

			mockResponse := [][]any{
				{"hello", "123"},
			}
			result := &Response{
				Data: mockResponse,
				ColumnDefinition: []ColumnInfo{
					{
						Name:        "abc",
						DataTypeOID: uint32(oid.T__text),
					},
					{
						Name:        "numbers",
						DataTypeOID: uint32(oid.T__numeric),
					},
				},
			}

			// Write the results back to the PostgreSQL client
			if err := writeResults(writer, result); err != nil {
				// log.Error("✏️ Write result failed", "error", err)
				return err
			}

			return writer.Complete("OK")
		}, wire.WithParameters(wire.ParseParameters(query)), wire.WithColumns(table)) // Injects parameter parsing logic

		return wire.Prepared(stmt), nil
	}
}

// writeResults takes the result from the HTTP response and returns it to the wire protocol client
func writeResults(writer wire.DataWriter, result *Response) error {
	columns := make(wire.Columns, len(result.ColumnDefinition))
	for i, info := range result.ColumnDefinition {
		columns[i] = wire.Column{
			Table: int32(i),
			Name:  info.Name,
			Oid:   oid.Oid(info.DataTypeOID),
		}
	}

	// // Send column definitions to client
	// if err := writer.Define(columns); err != nil {
	// 	return fmt.Errorf("define columns failed: %w", err)
	// }

	// Write each row back
	for _, row := range result.Data {
		if err := writer.Row(row); err != nil {
			return fmt.Errorf("write row failed: %w", err)
		}
	}

	return nil
}

# poc-prisma-psql-wire

## Type of PoC

### 1. Prisma app to postgresql directly

In folder `prisma_app`, just run `docker-compose up --build`

### 2. psql wire that got hit by psql client

In folder `psql_wire_app`, just run `docker-compose up --build`

### 3. Prisma app with psql wire

In root folder, just run `docker-compose up --build`


## Questions:

1. Before try to remake this poc, prisma feels like cannot connect to psql, but now the error change to 
```meta: {
prisma_app-1        |     code: 'XXUUU',
prisma_app-1        |     message: 'ERROR: write row failed: unable to encode "hello" into binary format for _text (OID 1009): cannot find encode plan'
prisma_app-1        |   }```

2. At line [93-96](https://github.com/Ace002/poc-prisma-psql-wire/blob/e39ccd9179c0ab44f04b1314a917fd856e389729/psql_wire_app/handlers/handler.go#L93) in older version, there is a define so that I can get the result first then construct the columns with this `Define` method, but with the latest one, I see that in a simple examples, it defines the table at first. Can I construct it when I have received and it will be depends on the query?
const { PrismaClient } = require('@prisma/client');

const prisma = new PrismaClient();

async function main() {
  try {
    // Attempt to connect and run a simple query
    const result = await prisma.$queryRaw`SELECT 1`;
    console.log('✅ PostgreSQL connection successful:', result);
  } catch (error) {
    console.error('❌ Failed to connect to PostgreSQL:', error);
    process.exit(1); // Exit with an error code
  } finally {
    await prisma.$disconnect();
  }
}

main();
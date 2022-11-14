require("dotenv/config");

// const {
//   PrismaClient: PrismaClientServer,
// } = require("./prisma-sqlserver/client");
const {
  PrismaClient: PrismaClientPostgres,
} = require("./prisma-postgres/client");

console.time();
const prisma = new PrismaClientPostgres();

prisma.$queryRaw`select nomcli, codcli from tbl_cliente`.then((data) => {
  console.log(data.length);
  console.timeEnd();
});

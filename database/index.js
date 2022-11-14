require("dotenv/config");

const { PrismaClient } = require("./prisma-sqlserver/client");

console.time();

const prisma = new PrismaClient();

prisma.$queryRaw`select TOP 1 nomcli, codcli from sapiens_teste.dbo.e085cli`.then(
  (data) => {
    console.log(data.length);
    console.timeEnd();
  }
);

FROM mcr.microsoft.com/mssql/server:2017-latest

WORKDIR /var/opt/mssql/

COPY migration.sql ./migration.sql

#RUN /opt/mssql-tools/bin/sqlcmd -U sa -P Qwerty1234 -i migration.sql
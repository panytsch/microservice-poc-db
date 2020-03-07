<h2>In docker</h2>

````
docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=Qwerty1234" -p 1433:1433 -v "%cd%/msdata/data:/var/opt/mssql/data" -v "%cd%/msdata/log:/var/opt/mssql/log" -v "%cd%/msdata/secrets:/var/opt/mssql/secrets" --name=sql1 -d mcr.microsoft.com/mssql/server:2019-GA-ubuntu-16.04
````
User = sa<br>
Password = Qwerty11234<br>
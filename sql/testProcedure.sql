CREATE PROCEDURE dbo.testProcedure @a bit
AS
    SET NOCOUNT ON;
    if @a = 1
        select one = 12,
               two = N'asdaaaa'
    else
        select ReturnCode = -5
GO

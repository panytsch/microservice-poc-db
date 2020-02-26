CREATE PROCEDURE dbo.twoDataSets
AS
    SET NOCOUNT ON;
    SELECT  one = 1,
            two = '2';
    
    SELECT  ReturnCode = 1
GO

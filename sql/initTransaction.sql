create procedure dbo.createTransaction @userId int, @status int, @amount money
as
    SET NOCOUNT ON
declare
    @transactionID int = 0;

    if EXISTS(select id
              from users
              where dbo.users.id = @userId)
        begin
            insert into transactions (status, user_id, amount) values (@status, @userId, @amount);
            set @transactionID = SCOPE_IDENTITY();
        end

select *
from transactions (nolock)
where id = @transactionID
    if @transactionID = 0
        select ReturnCode = -1;
    else
        select ReturnCode = 1;
go


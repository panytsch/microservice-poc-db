create procedure dbo.createPayment @userId int, @status int, @amount money
as
    SET NOCOUNT ON
declare
    @paymentID int = 0;

    if EXISTS(select id
              from users
              where dbo.users.id = @userId)
        begin
            insert into payments (status, user_id, amount) values (@status, @userId, @amount);
            set @paymentID = SCOPE_IDENTITY();
        end

select *
from payments (nolock)
where id = @paymentID
    if @paymentID = 0
        select ReturnCode = -1;
    else
        select ReturnCode = 1;
go


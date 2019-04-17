USE `db_intern`;
DROP procedure IF EXISTS `spCreateUser`;

DELIMITER $$
USE `db_intern`$$
CREATE PROCEDURE `spCreateUser` (
IN p_Username varchar(25),
IN p_EmailId varchar(50),
IN p_PhoneNo varchar(10),
IN p_Password varchar(50)
)
BEGIN

if ( select exists (select 1 from userData where emailId = p_EmailId) ) THEN
UPDATE userData SET 
    UserName=p_Username,
    phoneNo=p_PhoneNo,
    password=p_Password,
    dateTime=NOW()
WHERE emailId = p_EmailId ;
ELSE

insert into userData
(
    UserName,
    emailId,
    phoneNo,
    password,
    dateTime
)
values
(
    p_Username,
    p_EmailId,
    p_PhoneNo,
    p_Password,
    NOW()
);

END IF;

END$$

DELIMITER ;

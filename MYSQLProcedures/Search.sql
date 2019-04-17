USE `db_intern`;
DROP procedure IF EXISTS `spSearch`;

DELIMITER $$
USE `db_intern`$$
CREATE PROCEDURE `spSearch` (
IN p_EmailId varchar(50)
)
BEGIN
select * from userData where emailId=p_EmailId;

END$$

DELIMITER ;

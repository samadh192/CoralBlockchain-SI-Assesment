USE `db_intern`;
DROP procedure IF EXISTS `spDelete`;

DELIMITER $$
USE `db_intern`$$
CREATE PROCEDURE `spDelete` (
IN p_EmailId varchar(50)
)
BEGIN
DELETE from userData where emailId=p_EmailId;

END$$

DELIMITER ;

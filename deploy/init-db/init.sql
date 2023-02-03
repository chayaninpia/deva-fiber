CREATE TABLE t_power (
  active_power INT,
  power_input INT
);

DELIMITER //
CREATE PROCEDURE insert_random_values()
BEGIN
  DECLARE i INT DEFAULT 1;
  WHILE i <= 1000 DO
    INSERT INTO t_power (active_power, power_input)
    VALUES (ROUND(RAND() * 1000), ROUND(RAND() * 1000));
    SET i = i + 1;
  END WHILE;
END//
DELIMITER ;

CALL insert_random_values();



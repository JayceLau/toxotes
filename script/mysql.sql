-- 创建数据库用户
CREATE USER 'toxotes'@'%' IDENTIFIED BY 'password';
-- 创建用户服务数据库
CREATE DATABASE toxot_user DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_bin;
-- 授权用户服务数据库给toxotes
GRANT ALL ON toxot_user.* TO 'toxotes'@'%';
FLUSH PRIVILEGES;

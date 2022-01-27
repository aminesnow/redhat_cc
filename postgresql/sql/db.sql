-- Clean db first to be able to remove user
DROP DATABASE IF EXISTS bucket_store;

-- Create common resources admin user
DROP USER IF EXISTS bucket_store_admin;
CREATE USER bucket_store_admin WITH CREATEDB CREATEROLE PASSWORD :'bucket_store_admin_pwd';

-- Create database
CREATE DATABASE bucket_store OWNER bucket_store_admin;
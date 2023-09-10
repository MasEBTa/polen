/*Tabel User_credential*/
CREATE TABLE
  public.user_credential
(
  id VARCHAR(225) PRIMARY KEY NOT NULL,
  username VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(225) NOT NULL,
  is_active BOOLEAN
);
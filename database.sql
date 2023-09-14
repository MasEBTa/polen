/*Tabel User_credential*/ -- fixed
CREATE TABLE user_credential
(
  id VARCHAR(225) PRIMARY KEY NOT NULL,
  username VARCHAR(100) NOT NULL UNIQUE,
  email VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(225) NOT NULL,
  role VARCHAR(50) NOT NULL,
  virtual_account_number VARCHAR(225) DEFAULT "",
  is_active BOOLEAN
);
-- fixed
INSERT INTO user_credential (id, username, email, password, role, is_active) VALUES ('456', 'admin', 'compani.mail.yo', '$2a$10$FTqRPKh1IrHzvzi1YbhTbOY0pk.zQPAnh7OxJxK7D4YEih2GG2DqK','admin', true);

-- Account Table
CREATE TABLE biodata (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    user_credential_id VARCHAR(55) NOT NULL,
    full_name VARCHAR(255),
    nik VARCHAR(20),
    phone_number VARCHAR(20),
    occupation VARCHAR(255),
    place_of_birth VARCHAR(255),
    date_of_birth DATE,
    postal_code VARCHAR(10),
    is_eglible BOOLEAN,
	status_update BOOLEAN,
	additional_information TEXT NULL DEFAULT 'biodata is not updated',
    FOREIGN KEY (user_credential_id) REFERENCES public.user_credential (id)
);

-- Account Table
CREATE TABLE saldo (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    user_credential_id VARCHAR(55) NOT NULL,
    total_saving INT,
    FOREIGN KEY (user_credential_id) REFERENCES public.user_credential (id)
);
-- fixed
INSERT INTO saldo (id, user_credential_id, total_saving) VALUES ('789', '456', 100000000);

-- Deposit Interest Table
CREATE TABLE deposit_interest (
    id SERIAL PRIMARY KEY NOT NULL,
    created_date DATE,
    interest_rate DECIMAL(5, 2)
);

-- Deposit Table
CREATE TABLE deposit (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    user_credential_id VARCHAR(55) NOT NULL,
    deposit_amount DECIMAL(15, 2),
    interest_rate DECIMAL(5, 2) NOT NULL,
    created_date DATE,
    maturity_date DATE,
    FOREIGN KEY (user_credential_id) REFERENCES user_credential (id)
);

-- Junction Table between Deposit and Deposit Interest (many-to-many)
-- CREATE TABLE deposit_deposit_interest (
--     deposit_id VARCHAR(55),
--     deposit_interest_id INT,
--     FOREIGN KEY (deposit_id) REFERENCES deposit (id),
--     FOREIGN KEY (deposit_interest_id) REFERENCES deposit_interest (id)
-- );

-- Top Up Table
CREATE TABLE top_up (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    user_credential_id VARCHAR(55) NOT NULL,
    top_up_amount DECIMAL(15, 2),
    countdown_time TIMESTAMP,
    status VARCHAR(20),
    FOREIGN KEY (user_credential_id) REFERENCES user_credential (id)
);

-- application cost Table
CREATE TABLE application_handling_cost (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    nominal INT NOT NULL,
    unit VARCHAR(100) NOT NULL UNIQUE,
);

-- -- Loan Duration Table
-- CREATE TABLE loan_duration (
--     id SERIAL PRIMARY KEY NOT NULL,
--     duration_months INT NOT NULL,
--     loan_interest_rate DECIMAL(5, 2) NOT NULL
-- );

-- Loan duration Table
CREATE TABLE loan_duration (
    id SERIAL PRIMARY KEY NOT NULL,
    duration_months INT NOT NULL,
    loan_interest_rate DECIMAL(5, 2) NOT NULL
);

-- Loan Table
CREATE TABLE loan (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    user_credential_id VARCHAR(55) NOT NULL,
    loan_amount DECIMAL(15, 2),
    loan_duration INT,
    loan_interest_rate DECIMAL(5, 2) NOT NULL,
    application_handling_cost_nominal INT NOT NULL,
    application_handling_cost_unit INT NOT NULL,
    total_amount_of_dept INT NOT NULL,
    loan_date_created DATE,
    status VARCHAR(20),
    FOREIGN KEY (user_credential_id) REFERENCES user_credential (id)
);

-- Installenment Loan
CREATE TABLE installenment_loan (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    user_credential_id VARCHAR(55) NOT NULL,
    loan_id VARCHAR(55) NOT NULL,
    isPayed BOOLEAN,
    payment_installenment INT,
    payment_deadline DATE,
    application_handling_cost_nominal INT NOT NULL,
    application_handling_cost_unit INT NOT NULL,
    total_amount_of_dept INT NOT NULL,
    loan_date_created DATE,
    status VARCHAR(20),
    FOREIGN KEY (user_credential_id) REFERENCES user_credential (id)
    FOREIGN KEY (loan_id) REFERENCES loan (id)
);

-- Loan Payment Table
CREATE TABLE loan_payment (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    loan_id VARCHAR(55) NOT NULL,
    payment_amount DECIMAL(15, 2),
    virtual_account_number VARCHAR(20),
    countdown_time TIMESTAMP,
    status VARCHAR(20),
    FOREIGN KEY (loan_id) REFERENCES loan (id)
);

-- -- Table for Active Loan Duration of Borrowers
-- CREATE TABLE active_loan_duration (
--     id VARCHAR(55) PRIMARY KEY NOT NULL,
--     user_credential_id VARCHAR(55) NOT NULL,
--     active_loan_count INT,
--     FOREIGN KEY (user_credential_id) REFERENCES user_credential (id)
-- );

-- Transaction Table
CREATE TABLE transaction (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    user_credential_id VARCHAR(55) NOT NULL,
    transaction_type VARCHAR(20) NOT NULL,
    transaction_amount DECIMAL(15, 2) NOT NULL,
    transaction_date DATE NOT NULL,
    FOREIGN KEY (user_credential_id) REFERENCES user_credential (id)
);

-- Loan History Table
CREATE TABLE loan_history (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    account_id VARCHAR(55) NOT NULL,
    loan_amount DECIMAL(15, 2) NOT NULL,
    loan_date DATE NOT NULL,
    loan_status VARCHAR(20) NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account (id)
);

-- Deposit History Table
-- CREATE TABLE deposit_history (
--     id VARCHAR(55) PRIMARY KEY NOT NULL,
--     account_id VARCHAR(55) NOT NULL,
--     modal_amount DECIMAL(15, 2) NOT NULL,
--     deposit_date DATE NOT NULL,
--     interest_rate DECIMAL(5, 2) NOT NULL,
--     status VARCHAR(20) NOT NULL,
--     FOREIGN KEY (account_id) REFERENCES account (id)
-- );

DROP TABLE deposit_history;
DROP TABLE loan_history;
DROP TABLE saldo;
DROP TABLE transaction;
DROP TABLE active_loan_duration;
DROP TABLE loan_payment;
DROP TABLE loan;
DROP TABLE loan_duration;
DROP TABLE top_up;
DROP TABLE deposit_deposit_interest;
DROP TABLE deposit;
DROP TABLE deposit_interest;
DROP TABLE account;
DROP TABLE user_credential;
/*Tabel User_credential*/
CREATE TABLE
  public.user_credential
(
  id VARCHAR(225) PRIMARY KEY NOT NULL,
  username VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(225) NOT NULL,
  role VARCHAR(50) NOT NULL UNIQUE,
  is_active BOOLEAN
);

-- Account Table
CREATE TABLE account (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    user_credential_id VARCHAR(55) NOT NULL,
    full_name VARCHAR(255),
    nik VARCHAR(20),
    phone_number VARCHAR(20),
    occupation VARCHAR(255),
    place_of_birth VARCHAR(255),
    date_of_birth DATE,
    postal_code VARCHAR(10),
    FOREIGN KEY (user_credential_id) REFERENCES public.user_credential (id)
);

-- Deposit Interest Table
CREATE TABLE deposit_interest (
    id SERIAL PRIMARY KEY NOT NULL,
    created_date DATE,
    interest_rate DECIMAL(5, 2)
);

-- Deposit Table
CREATE TABLE deposit (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    account_id VARCHAR(55) NOT NULL,
    deposit_amount DECIMAL(15, 2),
    deposit_interest_id INT,
    created_date DATE,
    maturity_date DATE,
    FOREIGN KEY (account_id) REFERENCES account (id)
);

-- Junction Table between Deposit and Deposit Interest (many-to-many)
CREATE TABLE deposit_deposit_interest (
    deposit_id VARCHAR(55),
    deposit_interest_id INT,
    FOREIGN KEY (deposit_id) REFERENCES deposit (id),
    FOREIGN KEY (deposit_interest_id) REFERENCES deposit_interest (id)
);

-- Top Up Table
CREATE TABLE top_up (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    account_id VARCHAR(55) NOT NULL,
    top_up_amount DECIMAL(15, 2),
    virtual_account_number VARCHAR(20),
    countdown_time TIMESTAMP,
    status VARCHAR(20),
    FOREIGN KEY (account_id) REFERENCES account (id)
);

-- Loan Duration Table
CREATE TABLE loan_duration (
    id SERIAL PRIMARY KEY NOT NULL,
    duration_months INT
);

-- Loan Table
CREATE TABLE loan (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    account_id VARCHAR(55) NOT NULL,
    loan_amount DECIMAL(15, 2),
    loan_duration_id INT,
    loan_date DATE,
    status VARCHAR(20),
    FOREIGN KEY (account_id) REFERENCES account (id),
    FOREIGN KEY (loan_duration_id) REFERENCES loan_duration (id)
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

-- Table for Active Loan Duration of Borrowers
CREATE TABLE active_loan_duration (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    account_id VARCHAR(55) NOT NULL,
    active_loan_count INT,
    FOREIGN KEY (account_id) REFERENCES account (id)
);

-- Transaction Table
CREATE TABLE transaction (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    account_id VARCHAR(55) NOT NULL,
    transaction_type VARCHAR(20) NOT NULL,
    transaction_amount DECIMAL(15, 2) NOT NULL,
    transaction_date DATE NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account (id)
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
CREATE TABLE deposit_history (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    account_id VARCHAR(55) NOT NULL,
    modal_amount DECIMAL(15, 2) NOT NULL,
    deposit_date DATE NOT NULL,
    interest_rate DECIMAL(5, 2) NOT NULL,
    status VARCHAR(20) NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account (id)
);

/*Tabel Biodata User*/
CREATE TABLE public.biodata_user
(
  id VARCHAR(225) PRIMARY KEY NOT NULL,
  user_credential_id VARCHAR(225) NOT NULL REFERENCES public.user_credential(id),
  nama_lengkap VARCHAR(225) NOT NULL,
  nik VARCHAR(20) NOT NULL UNIQUE,
  nomor_telepon VARCHAR(20) NOT NULL UNIQUE,
  pekerjaan VARCHAR(225) NOT NULL,
  tempat_lahir VARCHAR(225) NOT NULL,
  tanggal_lahir DATE NOT NULL,
  kode_pos VARCHAR(10) NOT NULL
);

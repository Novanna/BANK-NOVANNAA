-- ----------------------------
-- Table structure for Customers
-- ----------------------------

CREATE TABLE IF NOT EXISTS customers (
      id SERIAL NOT NULL,
      nama_lengkap varchar(255) DEFAULT NULL,
      alamat varchar(450) DEFAULT NULL,
      tanggal_lahir varchar(255) DEFAULT NULL,
      tempat_lahir varchar(255) DEFAULT NULL,
      jenis_kelamin varchar(255) DEFAULT NULL,
      no_ktp BIGINT UNIQUE NULL,
      no_hp BIGINT DEFAULT NULL, 
      created_at TIMESTAMPTZ DEFAULT NOW(),
      updated_at TIMESTAMPTZ DEFAULT NOW(),
      PRIMARY KEY (id)
);

-- ----------------------------
-- Table structure for Admin
-- ----------------------------

CREATE TABLE IF NOT EXISTS admins ( 
      id SERIAL NOT NULL,
      first_name varchar(255) DEFAULT NULL,
      last_name varchar(255) DEFAULT NULL,
      email varchar(255) DEFAULT NULL,
      password varchar(255) DEFAULT NULL,
      created_at TIMESTAMPTZ DEFAULT NOW(),
      updated_at TIMESTAMPTZ DEFAULT NOW(),
      PRIMARY KEY (id)
);
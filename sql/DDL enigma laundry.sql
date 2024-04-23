CREATE DATABASE enigma_laundry;

CREATE TABLE mst_customer (
	id INT NOT NULL,
	nama VARCHAR(255) NOT NULL,
	no_hp VARCHAR(50) NOT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE mst_layanan (
	id INT NOT NULL,
	nama_layanan VARCHAR(100) NOT NULL,
	satuan VARCHAR(100) NOT NULL,
	harga INT NOT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE trx_laundry (
	id INT NOT NULL PRIMARY KEY,
	no_trx VARCHAR(50) NOT NULL,
	customer_id INT NOT NULL,
	tanggal_masuk DATE NOT NULL,
	tanggal_selesai DATE NOT NULL,
	penerima VARCHAR(100) NOT NULL,
	total_biaya INT,
	FOREIGN KEY(customer_id) REFERENCES mst_customer(id)
);

CREATE TABLE trx_laundry_detail (
	id INT NOT NULL PRIMARY KEY,
	trx_laundry_id INT NOT NULL,
	layanan_id INT NOT NULL,
	jumlah INT NOT NULL,
	FOREIGN KEY(trx_laundry_id) REFERENCES trx_laundry(id),
	FOREIGN KEY(layanan_id) REFERENCES mst_layanan(id)
);


SELECT * FROM mst_customer;
SELECT * FROM mst_layanan;
SELECT * FROM trx_laundry;
SELECT * FROM trx_laundry_detail;

INSERT INTO mst_customer (id, nama, no_hp) 
VALUES (1, 'TERIZLA', '082121212121');
INSERT INTO mst_customer (id, nama, no_hp) 
VALUES (2, 'CHOUNIMA', '083131313131');
INSERT INTO mst_customer (id, nama, no_hp) 
VALUES (3, 'YUZHONG', '084141414141');
INSERT INTO mst_customer (id, nama, no_hp) 
VALUES (4, 'DYROT', '085151515151');
INSERT INTO mst_customer (id, nama, no_hp) 
VALUES (5, 'LAPU-LAPU', '086161616161');


INSERT INTO mst_layanan (id, nama_layanan, satuan, harga) 
VALUES (1, 'Cuci + Setrika', 'KG', 7000);
INSERT INTO mst_layanan (id, nama_layanan, satuan, harga) 
VALUES (2, 'Laundry bedcover', 'Buah', 50000);
INSERT INTO mst_layanan (id, nama_layanan, satuan, harga) 
VALUES (3, 'Laundry bonek', 'Buah', 25000);


INSERT INTO trx_laundry (id, no_trx, customer_id, tanggal_masuk, tanggal_selesai,
						 penerima)
VALUES (1, 'trx01', 2, '2024-04-22', '2024-04-26', 'Butet');
INSERT INTO trx_laundry (id, no_trx, customer_id, tanggal_masuk, tanggal_selesai,
						 penerima)
VALUES (2, 'trx02', 3, '2024-04-28', '2024-04-30', 'Butet');
UPDATE trx_laundry SET total_biaya = 0;


INSERT INTO trx_laundry_detail (id, trx_laundry_id, layanan_id, jumlah)
VALUES (1, 1, 1, 5);
INSERT INTO trx_laundry_detail (id, trx_laundry_id, layanan_id, jumlah)
VALUES (2, 1, 2, 2);
INSERT INTO trx_laundry_detail (id, trx_laundry_id, layanan_id, jumlah)
VALUES (3, 2, 3, 3);
INSERT INTO trx_laundry_detail (id, trx_laundry_id, layanan_id, jumlah)
VALUES (4, 2, 1, 2);

UPDATE trx_laundry SET total_biaya = 135000 WHERE id = 1;
UPDATE trx_laundry SET total_biaya = 89000 WHERE id = 2;
package main

import (
	"bufio"
	"database/sql"
	"enigma_laundry/entity"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "enigma_laundry"
)

// pre define db
var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	for {
		fmt.Println(strings.Repeat("=", 30))
		fmt.Print(strings.Repeat("=", 7))
		fmt.Print(" Enigma Laundry ")
		fmt.Print(strings.Repeat("=", 7))
		fmt.Println()
		fmt.Println("List menu laundry :")
		fmt.Println("1. Customer Menu")
		fmt.Println("2. Layanan Menu")
		fmt.Println("3. Transaksi Menu")
		fmt.Println("4. Exit")
		fmt.Println(strings.Repeat("=", 30))

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter your choice : ")
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			for {
				fmt.Println(strings.Repeat("=", 30))
				fmt.Print(strings.Repeat("=", 7))
				fmt.Print(" Customer Menu ")
				fmt.Print(strings.Repeat("=", 8))
				fmt.Println()
				fmt.Println("1. View All Customer")
				fmt.Println("2. View Customer Detail")
				fmt.Println("3. Add New Customer")
				fmt.Println("4. Update Customer")
				fmt.Println("5. Delete Customer")
				fmt.Println("6. Back")
				fmt.Println(strings.Repeat("=", 30))

				fmt.Print("Enter your choice : ")
				scanner.Scan()
				choiceCustomer, _ := strconv.Atoi(scanner.Text())

				if choiceCustomer == 1 {
					fmt.Println("All Customer")
					customers := getAllCustomers()
					for _, customer := range customers {
						fmt.Println(customer.Id, customer.Nama, customer.No_hp)
					}
				} else if choiceCustomer == 2 {
					fmt.Print("Enter Customer ID : ")
					scanner.Scan()
					customerId, _ := strconv.Atoi(scanner.Text())

					customer, err := getCustomerById(customerId)
					if err != nil {
						fmt.Println("Customer not found")
						continue
					} else {
						fmt.Println("Customer Detail")
						fmt.Println("Nama:", customer.Nama)
						fmt.Println("No Hp:", customer.No_hp)
					}

				} else if choiceCustomer == 3 {
					customer := entity.Customer{}
					fmt.Println("Input Data customer")
					fmt.Print("Customer ID : ")
					scanner.Scan()
					customer.Id, _ = strconv.Atoi(scanner.Text())
					if customer.Id == 0 {
						fmt.Println("Customer ID harus diisi")
						continue
					}

					fmt.Print("Nama Customer : ")
					scanner.Scan()
					customer.Nama = scanner.Text()
					if customer.Nama == "" {
						fmt.Println("Nama tidak boleh kosong")
						continue
					}

					fmt.Print("No Hp : ")
					scanner.Scan()
					customer.No_hp = scanner.Text()
					if customer.No_hp == "" {
						fmt.Println("No hp tidak boleh kosong")
						continue
					}

					addCustomer(customer)

				} else if choiceCustomer == 4 {
					fmt.Println("Update Customer")
					fmt.Print("Enter Customer ID : ")
					scanner.Scan()
					id, _ := strconv.Atoi(scanner.Text())

					customerDB, err := getCustomerById(id)
					if err != nil {
						fmt.Println("Customer not found")
						continue
					} else {
						customer := entity.Customer{}

						fmt.Println("Update Detail Customer! ")
						fmt.Print("Update Nama : ")
						scanner.Scan()
						customer.Nama = scanner.Text()
						fmt.Print("Update No Hp : ")
						scanner.Scan()
						customer.No_hp = scanner.Text()

						if customer.Nama == "" {
							customer.Nama = customerDB.Nama
						} else if customer.No_hp == "" {
							customer.No_hp = customerDB.No_hp
						}

						fmt.Print("Yakin ingin mengupdate data ini ? (Y/N) : ")
						scanner.Scan()
						confirm := scanner.Text()

						if strings.ToLower(confirm) == "y" {
							updateCustomer(id, customer)
						} else if strings.ToLower(confirm) == "n" {
							fmt.Println("Update dibatalkan")
						} else {
							fmt.Println("Input invalid")
						}
					}

				} else if choiceCustomer == 5 {
					fmt.Print("Enter Customer ID : ")
					scanner.Scan()
					id, _ := strconv.Atoi(scanner.Text())

					fmt.Print("Yakin ingin menghapus data ini ? (Y/N) : ")
					scanner.Scan()
					confirm := scanner.Text()

					if strings.ToLower(confirm) == "y" {
						deleteCustomer(id)
					} else if strings.ToLower(confirm) == "n" {
						fmt.Println("Data batal di hapus")
					} else {
						fmt.Println("Input invalid")
					}

				} else if choiceCustomer == 6 {
					break
				} else {
					fmt.Println("Input invalid")
				}
			}
		case 2:
			for {
				fmt.Println(strings.Repeat("=", 30))
				fmt.Print(strings.Repeat("=", 7))
				fmt.Print(" Layanan Menu ")
				fmt.Print(strings.Repeat("=", 9))
				fmt.Println()
				fmt.Println("1. View All Layanan")
				fmt.Println("2. View Layanan Detail")
				fmt.Println("3. Add New Layanan")
				fmt.Println("4. Update Layanan")
				fmt.Println("5. Delete Layanan")
				fmt.Println("6. Back")
				fmt.Println(strings.Repeat("=", 30))

				fmt.Print("Enter your choice : ")
				scanner.Scan()
				choiceLayanan, _ := strconv.Atoi(scanner.Text())

				if choiceLayanan == 1 {
					fmt.Println("All Layanan")
					layanan := getAllLayanan()
					for _, value := range layanan {
						fmt.Println(value.Id, value.NamaLayanan, value.Satuan, value.Harga)
					}
				} else if choiceLayanan == 2 {
					fmt.Print("Enter Layanan ID : ")
					scanner.Scan()
					layananId, _ := strconv.Atoi(scanner.Text())

					Layanan, err := getLayananById(layananId)
					if err != nil {
						fmt.Println("Layanan not found")
					} else {
						fmt.Println("Layanan Detail")
						fmt.Println("Nama Layanan:", Layanan.NamaLayanan)
						fmt.Println("Satuan:", Layanan.Satuan)
						fmt.Println("Harga:", Layanan.Harga)
					}

				} else if choiceLayanan == 3 {
					Layanan := entity.Layanan{}
					fmt.Println("Input Data Layanan")
					fmt.Print("Layanan ID : ")
					scanner.Scan()
					Layanan.Id, _ = strconv.Atoi(scanner.Text())
					if Layanan.Id == 0 {
						fmt.Println("Layanan ID harus diisi")
						continue
					}

					fmt.Print("Nama Layanan : ")
					scanner.Scan()
					Layanan.NamaLayanan = scanner.Text()
					if Layanan.NamaLayanan == "" {
						fmt.Println("Nama Layanan tidak boleh kosong")
						continue
					}

					fmt.Print("Satuan : ")
					scanner.Scan()
					Layanan.Satuan = scanner.Text()
					if Layanan.Satuan == "" {
						fmt.Println("Satuan tidak boleh kosong")
						continue
					}

					fmt.Print("Harga : ")
					scanner.Scan()
					Layanan.Harga, _ = strconv.Atoi(scanner.Text())
					if Layanan.Harga == 0 {
						fmt.Println("Harga tidak boleh kosong")
						continue
					}

					addLayanan(Layanan)

				} else if choiceLayanan == 4 {
					fmt.Println("Update Layanan")
					fmt.Print("Enter Layanan ID : ")
					scanner.Scan()
					id, _ := strconv.Atoi(scanner.Text())

					layananDB, err := getLayananById(id)
					if err != nil {
						fmt.Println("Layanan not found")
					} else {

						layanan := entity.Layanan{}

						fmt.Println("Update Detail Layanan! ")
						fmt.Print("Nama Layanan : ")
						scanner.Scan()
						layanan.NamaLayanan = scanner.Text()
						fmt.Print("Satuan : ")
						scanner.Scan()
						layanan.Satuan = scanner.Text()
						fmt.Print("Harga Layanan : ")
						scanner.Scan()
						layanan.Harga, _ = strconv.Atoi(scanner.Text())

						if layanan.NamaLayanan == "" {
							layanan.NamaLayanan = layananDB.NamaLayanan
						}
						if layanan.Satuan == "" {
							layanan.Satuan = layananDB.Satuan
						}
						if layanan.Harga == 0 {
							layanan.Harga = layananDB.Harga
						}

						fmt.Print("Yakin ingin mengupdate data ini ? (Y/N) : ")
						scanner.Scan()
						confirm := scanner.Text()
						if strings.ToLower(confirm) == "y" {
							updateLayanan(id, layanan)
						} else if strings.ToLower(confirm) == "n" {
							fmt.Println("Update dibatalkan")
						} else {
							fmt.Println("Update dibatalkan")
						}

					}

				} else if choiceLayanan == 5 {
					fmt.Print("Enter Layanan ID : ")
					scanner.Scan()
					id, _ := strconv.Atoi(scanner.Text())

					fmt.Print("Yakin ingin menghapus data ini ? (Y/N) : ")
					scanner.Scan()
					confirm := scanner.Text()

					if strings.ToLower(confirm) == "y" {
						deleteLayanan(id)
					} else if strings.ToLower(confirm) == "n" {
						fmt.Println("Dibatalkan")
					} else {
						fmt.Println("Input invalid")
					}

				} else if choiceLayanan == 6 {
					break
				} else {
					fmt.Println("Input invalid")
				}
			}
		case 3:
			for {
				fmt.Println(strings.Repeat("=", 30))
				fmt.Print(strings.Repeat("=", 6))
				fmt.Print(" Transaction Menu ")
				fmt.Print(strings.Repeat("=", 6))
				fmt.Println()
				fmt.Println("1. View All Transaction")
				fmt.Println("2. View Transaction Detail")
				fmt.Println("3. Add New Transaction")
				fmt.Println("4. Back")
				fmt.Println(strings.Repeat("=", 30))

				fmt.Print("Enter your choice : ")
				scanner.Scan()
				choiceTrx, _ := strconv.Atoi(scanner.Text())

				if choiceTrx == 1 {
					fmt.Println("View All Transaction")
					trx := getAllTrx()

					fmt.Println("id|Notrx|customer_id|tanggal_masuk|tanggal_selesai|penerima|total_biaya|")
					for _, val := range trx {
						tanggalMasukStr := val.Tanggal_masuk.Format("2006-01-02")
						tanggalSelesaiStr := val.Tanggal_selesai.Format("2006-01-02") // Or your preferred format

						fmt.Println(val.Id, "|", val.NoTrx, "|  ", val.Customer_id, "    |", tanggalMasukStr, " |", tanggalSelesaiStr, "   |", val.Penerima, " |", val.Total_biaya)
					}
				} else if choiceTrx == 2 {
					fmt.Println("View Transaction Detail")
					fmt.Print("Enter Transaction id: ")
					scanner.Scan()
					id, _ := strconv.Atoi(scanner.Text())
					trxDetail := getAllTrxDetail(id)

					fmt.Println("id|trx_id|layanan_id|jml|")
					for _, val := range trxDetail {
						fmt.Println(val.Id, "|   ", val.Trx_laundry_id, "|       ", val.Layanan_id, "|", val.Jumlah, "|")
					}
				} else if choiceTrx == 3 {
					fmt.Println("Add New Transaction")
					fmt.Print("Enter Transaction id : ")
					scanner.Scan()
					id, _ := strconv.Atoi(scanner.Text())
					transaction(id)
				} else if choiceTrx == 4 {
					break
				} else {
					fmt.Println("Input invalid")
				}
			}
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Input invalid, silakan pilih ulang")
		}
	}

}

func koneksi_db() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func transaction(id int) {
	db := koneksi_db()
	defer db.Close()

	_, errr := findTrxById(id)

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	if errr != nil {
		createNewTrx(id, tx)
		createDetailTrx(id, tx)
		totalbiaya := getSumTotalBiaya(id, tx)
		updateTrx(id, totalbiaya, tx)
		err = tx.Commit()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Transaction Committed!")
		}
	} else {
		createDetailTrx(id, tx)
		totalbiaya := getSumTotalBiaya(id, tx)
		updateTrx(id, totalbiaya, tx)
		err = tx.Commit()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Transaction Committed!")
		}
	}
}

// transaction function
func getAllTrx() []entity.TrxLaundry {
	db := koneksi_db()
	defer db.Close()

	sqlStatement := "SELECT * FROM trx_laundry ORDER BY id ASC;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	trx := scanTrx(rows)

	return trx
}

func scanTrx(rows *sql.Rows) []entity.TrxLaundry {
	transaksi := []entity.TrxLaundry{}
	var err error

	for rows.Next() {
		trx := entity.TrxLaundry{}
		err := rows.Scan(&trx.Id, &trx.NoTrx, &trx.Customer_id, &trx.Tanggal_masuk, &trx.Tanggal_selesai, &trx.Penerima, &trx.Total_biaya)
		if err != nil {
			panic(err)
		}

		transaksi = append(transaksi, trx)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return transaksi
}

func getAllTrxDetail(trxId int) []entity.TrxLaundryDetail {
	db := koneksi_db()
	defer db.Close()

	sqlStatement := "SELECT * FROM trx_laundry_detail WHERE trx_laundry_id = $1;"

	rows, err := db.Query(sqlStatement, trxId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	trxDetails := scanTrxDetail(rows)

	return trxDetails
}

func scanTrxDetail(rows *sql.Rows) []entity.TrxLaundryDetail {
	trxDetails := []entity.TrxLaundryDetail{}
	var err error

	for rows.Next() {
		trxDetail := entity.TrxLaundryDetail{}
		err := rows.Scan(&trxDetail.Id, &trxDetail.Trx_laundry_id, &trxDetail.Layanan_id, &trxDetail.Jumlah)
		if err != nil {
			panic(err)
		}

		trxDetails = append(trxDetails, trxDetail)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return trxDetails
}

func updateTrx(trxId int, totalBiaya int, tx *sql.Tx) {
	sqlStatement := "UPDATE trx_laundry SET total_biaya = $2 WHERE id = $1;"

	_, err := tx.Exec(sqlStatement, trxId, totalBiaya)
	validate(err, "Update", tx)
}

func getSumTotalBiaya(id int, tx *sql.Tx) int {
	sqlStatement := "SELECT SUM(layanan.harga * trxd.jumlah) FROM trx_laundry_detail AS trxd INNER JOIN mst_layanan AS layanan ON trxd.layanan_id = layanan.id WHERE trx_laundry_id = $1;"

	totalBiaya := 0
	err := tx.QueryRow(sqlStatement, id).Scan(&totalBiaya)
	validate(err, "Sum", tx)

	return totalBiaya
}

func createNewTrx(id int, tx *sql.Tx) {
	newTransaksi := entity.TrxLaundry{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Create New Transaction")

	newTransaksi.Id = id

	fmt.Print("Nomor Transaksi : ")
	scanner.Scan()
	newTransaksi.NoTrx = scanner.Text()

	fmt.Print("Customer ID : ")
	scanner.Scan()
	newTransaksi.Customer_id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Tanggal masuk : ")
	scanner.Scan()
	tanggalMasukStr := scanner.Text()
	tanggalMasuk, _ := time.Parse("2006-01-02", tanggalMasukStr)
	newTransaksi.Tanggal_masuk = time.Date(tanggalMasuk.Year(), tanggalMasuk.Month(), tanggalMasuk.Day(), 0, 0, 0, 0, time.Local)

	fmt.Print("Tanggal selesai : ")
	scanner.Scan()
	tanggalSelesaiStr := scanner.Text()
	tanggalSelesai, _ := time.Parse("2006-01-02", tanggalSelesaiStr)
	newTransaksi.Tanggal_selesai = time.Date(tanggalSelesai.Year(), tanggalSelesai.Month(), tanggalSelesai.Day(), 0, 0, 0, 0, time.Local)

	fmt.Print("Penerima : ")
	scanner.Scan()
	newTransaksi.Penerima = scanner.Text()

	newTransaksi.Total_biaya = 0

	sqlStatement := "INSERT INTO trx_laundry (id, no_trx, customer_id, tanggal_masuk, tanggal_selesai, penerima, total_biaya) VALUES($1, $2, $3, $4, $5, $6, $7);"

	_, err := tx.Exec(sqlStatement, newTransaksi.Id, newTransaksi.NoTrx, newTransaksi.Customer_id, newTransaksi.Tanggal_masuk, newTransaksi.Tanggal_selesai, newTransaksi.Penerima, newTransaksi.Total_biaya)

	validate(err, "Create", tx)
}

func createDetailTrx(id int, tx *sql.Tx) {
	trxdetail := entity.TrxLaundryDetail{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Create Transaction Detail")
	fmt.Print("ID: ")
	scanner.Scan()
	trxdetail.Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Jenis layanan id : ")
	scanner.Scan()
	trxdetail.Layanan_id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Jumlah : ")
	scanner.Scan()
	trxdetail.Jumlah, _ = strconv.Atoi(scanner.Text())

	trxdetail.Trx_laundry_id = id

	sqlStatement := "INSERT INTO trx_laundry_detail (id, trx_laundry_id, layanan_id, jumlah) VALUES($1, $2, $3, $4);"

	_, err := tx.Exec(sqlStatement, trxdetail.Id, trxdetail.Trx_laundry_id, trxdetail.Layanan_id, trxdetail.Jumlah)

	validate(err, "Create", tx)
}

func findTrxById(id int) (entity.TrxLaundry, error) {
	db := koneksi_db()
	defer db.Close()
	var err error

	sqlStatement := "SELECT * FROM trx_laundry WHERE id = $1;"

	transaction := entity.TrxLaundry{}
	err = db.QueryRow(sqlStatement, id).Scan(&transaction.Id, &transaction.NoTrx, &transaction.Customer_id, &transaction.Tanggal_masuk, &transaction.Tanggal_selesai, &transaction.Penerima, &transaction.Total_biaya)

	return transaction, err
}

func validate(err error, msg string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println("Transaction rollback")
	} else {
		fmt.Println("Success " + msg + " Transaction")
	}
}

// Customer function
func deleteCustomer(id int) {
	db := koneksi_db()
	defer db.Close()
	var err error

	sqlStatement := "DELETE FROM mst_customer WHERE id = $1;"

	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Customer Berhasil dihapus")
	}
}
func updateCustomer(id int, customer entity.Customer) {
	db := koneksi_db()
	defer db.Close()
	var err error

	sqlStatement := "UPDATE mst_customer SET nama = $2, no_hp = $3 WHERE id = $1;"

	_, err = db.Exec(sqlStatement, id, customer.Nama, customer.No_hp)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Customer Berhasil diupdate")
	}
}
func addCustomer(customer entity.Customer) {
	db := koneksi_db()
	defer db.Close()
	var err error

	sqlStatement := "INSERT INTO mst_customer (id, nama, no_hp) VALUES($1, $2, $3);"

	_, err = db.Exec(sqlStatement, customer.Id, customer.Nama, customer.No_hp)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Customer Baru Berhasil ditambahkan")
	}
}
func getCustomerById(id int) (entity.Customer, error) {
	db := koneksi_db()
	defer db.Close()
	var err error

	sqlStatement := "SELECT * FROM mst_customer WHERE id = $1;"

	customer := entity.Customer{}
	err = db.QueryRow(sqlStatement, id).Scan(&customer.Id, &customer.Nama, &customer.No_hp)

	return customer, err
}
func getAllCustomers() []entity.Customer {
	db := koneksi_db()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_customer ORDER BY id ASC;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := scanCustomers(rows)

	return customers
}
func scanCustomers(rows *sql.Rows) []entity.Customer {
	customers := []entity.Customer{}
	var err error

	for rows.Next() {
		customer := entity.Customer{}
		err := rows.Scan(&customer.Id, &customer.Nama, &customer.No_hp)
		if err != nil {
			panic(err)
		}

		customers = append(customers, customer)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return customers
}

// Layanan function
func deleteLayanan(id int) {
	db := koneksi_db()
	defer db.Close()
	var err error

	sqlStatement := "DELETE FROM mst_layanan WHERE id = $1;"

	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Layanan Berhasil dihapus")
	}
}
func updateLayanan(id int, layanan entity.Layanan) {
	db := koneksi_db()
	defer db.Close()
	var err error

	sqlStatement := "UPDATE mst_layanan SET nama_layanan = $2, satuan = $3, harga = $4 WHERE id = $1;"

	_, err = db.Exec(sqlStatement, id, layanan.NamaLayanan, layanan.Satuan, layanan.Harga)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Layanan Berhasil diupdate")
	}
}
func addLayanan(layanan entity.Layanan) {
	db := koneksi_db()
	defer db.Close()
	var err error

	sqlStatement := "INSERT INTO mst_layanan (id, nama_layanan, satuan, harga) VALUES($1, $2, $3,$4);"

	_, err = db.Exec(sqlStatement, layanan.Id, layanan.NamaLayanan, layanan.Satuan, layanan.Harga)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Layanan Baru Berhasil ditambahkan")
	}
}
func getLayananById(id int) (entity.Layanan, error) {
	db := koneksi_db()
	defer db.Close()
	var err error

	sqlStatement := "SELECT * FROM mst_layanan WHERE id = $1;"

	layanan := entity.Layanan{}
	err = db.QueryRow(sqlStatement, id).Scan(&layanan.Id, &layanan.NamaLayanan, &layanan.Satuan, &layanan.Harga)

	return layanan, err
}
func getAllLayanan() []entity.Layanan {
	db := koneksi_db()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_layanan ORDER BY id ASC;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	layanan := scanLayanan(rows)

	return layanan
}
func scanLayanan(rows *sql.Rows) []entity.Layanan {
	layanans := []entity.Layanan{}
	var err error

	for rows.Next() {
		layanan := entity.Layanan{}
		err := rows.Scan(&layanan.Id, &layanan.NamaLayanan, &layanan.Satuan, &layanan.Harga)
		if err != nil {
			panic(err)
		}

		layanans = append(layanans, layanan)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return layanans
}

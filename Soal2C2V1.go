package main
import "fmt"
func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah : ")
	fmt.Scanln(&nam)
	if nam > 80 {
		//nam = " A "
	}

	if nam > 72.5 {
		//nam = " AB "
	}

	if nam > 65 {
		//nam = " B "
	}

	if nam > 57.5 {
		//nam = " BC "
	}

	if nam > 50 {
		//nam = " C "
	}

	if nam > 40 {
		//nam = " D "
	} else if nam <= 40 {
		//nam = " E "
	}
	fmt.Println("Nilai mata kuliah : ", nmk)
}
	
//A. Keluaran program ketika nam = 80.1 -> error (Tidak bisa menggunakan fungsi percabangan(NMK) seperti " A ", " AB ", " B ", dan sebagainya)

/*B. Kesalahan dalam Program :
1. Variabel nam dideklarasikan sebagai float64 tetapi malah diisi nilai string seperti " A ", " AB ", " B ", dan sebagainya sehingga program tidak dapat berjalan
2. penggunaan beberapa if statements tanpa else if menyebabkan semua kondisi diperiksa sehingga semua kondisi diperiksa dan nilai nam dapat diubah serta nilai nam dapat jatuh ke beberapa kategori
3. Program deklarasikan serta mencetak variabel nmk padahal variabel tersebut tidak memiliki nilai apapun*/

/*Alur program seharusnya : 
1. Deklarasikan variabel nam sebagai float64 untuk menyimpan inputan nilai user dan variabel nmk sebagai string untuk menyimpan nilai grade berupa huruf
2. Gunakan struktur if-else if untuk memeriksa nilai nam yang diinput user dan menetapkan nilai nmk yang sesuai dengan cara urutan kondisi / grade harus sesuai dengan rentang nilai yang diberikan tabel
3. Output yang akan ditampilkan adalah hasil variabel nmk yang berisi grade / kondisi yang sesuai*/
package request

type Product struct {
	Name  string `form:"name" valid:"required~field nama tidak ditemukan"`
	Price string `form:"price" valid:"required~field harga tidak ditemukan"`
	Stock string `form:"stock" valid:"required~field stok tidak ditemukan"`
}

type Activity struct {
	Group       uint   `form:"group" valid:"required~field kelompok kerja tidak ditemukan"`
	Title       string `form:"title" valid:"required~field judul tidak ditemukan"`
	Description string `form:"description" valid:"required~field deskripsi tidak ditemukan"`
}

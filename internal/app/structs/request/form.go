package request

type Product struct {
	Name  string `form:"name" valid:"required~field nama tidak ditemukan"`
	Price string `form:"price" valid:"required~field harga tidak ditemukan"`
	Stock string `form:"stock" valid:"required~field stok tidak ditemukan"`
}

package request

type SignIn struct {
	Username string `json:"username" valid:"required~field username tidak ditemukan"`
	Password string `json:"password" valid:"required~field password tidak ditemukan"`
}

type RegisterGuest struct {
	Fullname        string `json:"fullname" valid:"required~field nama lengkap tidak ditemukan"`
	Address         string `json:"address" valid:"required~field alamat tidak ditemukan"`
	PhoneNumber     string `json:"phone_number" valid:"required~field nomor handphone tidak ditemukan"`
	Username        string `json:"username" valid:"required~field username tidak ditemukan"`
	Password        string `json:"password" valid:"required~field password tidak ditemukan"`
	ConfirmPassword string `json:"confirm_password" valid:"required~field konfirmasi password tidak ditemukan"`
}

type Activation struct {
	Status bool `json:"status"`
}

type Category struct {
	Name string `json:"name" valid:"required~field nama tidak ditemukan"`
}

type Shop struct {
	CategoryUuid string `json:"category_uuid" valid:"required~field kategori tidak ditemukan"`
	Name         string `json:"name" valid:"required~field nama tidak ditemukan"`
	Owner        string `json:"owner" valid:"required~field nama pemilik tidak ditemukan"`
	Address      string `json:"address" valid:"required~field alamat tidak ditemukan"`
	PhoneNumber  string `json:"phone_number" valid:"required~field nomor handphone tidak ditemukan"`
}

type Product struct {
	Name  string `json:"name" valid:"required~field nama tidak ditemukan"`
	Price int64  `json:"price" valid:"required~field harga tidak ditemukan"`
	Stock int64  `json:"stock" valid:"required~field stok tidak ditemukan"`
}

type BuyProduct struct {
	ProductUuid string `json:"product_uuid" valid:"required~field Uuid produk tidak ditemukan"`
	Quantity    int64  `json:"quantity" valid:"required~field jumlah tidak ditemukan"`
}

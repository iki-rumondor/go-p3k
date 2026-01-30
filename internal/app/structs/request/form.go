package request

type Product struct {
	CategoryUuid string `form:"category_uuid" valid:"required~field kategori tidak ditemukan"`
	Name         string `form:"name" valid:"required~field nama tidak ditemukan"`
	Price        string `form:"price" valid:"required~field harga tidak ditemukan"`
	Unit         string `form:"unit" valid:"required~field satuan tidak ditemukan"`
	Stock        string `form:"stock" valid:"required~field stok tidak ditemukan"`
}

type Activity struct {
	Group       uint   `form:"group" valid:"required~field kelompok kerja tidak ditemukan"`
	Title       string `form:"title" valid:"required~field judul tidak ditemukan"`
	Description string `form:"description" valid:"required~field deskripsi tidak ditemukan"`
	Location    string `form:"location" valid:"required~field lokasi tidak ditemukan"`
	StartTime   int64  `form:"start_time" valid:"required~field waktu mulai tidak ditemukan"`
	EndTime     int64  `form:"end_time" valid:"required~field waktu selesai tidak ditemukan"`
}

type RegisterShop struct {
	RoleID          uint   `form:"role_id" valid:"required~field role pengguna tidak ditemukan"`
	ShopName        string `form:"shop_name" valid:"required~field nama toko tidak ditemukan"`
	Owner           string `form:"owner_name" valid:"required~field nama pemilik tidak ditemukan"`
	Address         string `form:"address" valid:"required~field alamat tidak ditemukan"`
	PhoneNumber     string `form:"phone_number" valid:"required~field nomor handphone tidak ditemukan"`
	Username        string `form:"username" valid:"required~field username tidak ditemukan"`
	Password        string `form:"password" valid:"required~field password tidak ditemukan"`
	ConfirmPassword string `form:"confirm_password" valid:"required~field konfirmasi password tidak ditemukan"`
}

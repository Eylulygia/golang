# Basit Kitap API

Kullanıcılar kayıt olur, giriş yapar, profilini görüntüler, kitap ödünç alır, iade eder ve isterse hesabını siler. Tüm akış bir MongoDB koleksiyonu üzerinde tutuluyor ve JWT ile güvene alınıyor.

## Neden Bu Proje?
- Go ile REST servisleri kurmayı öğrenmek.
- Fiber kullanarak basit ama düzenli bir router yapısı kurmak.
- MongoDB ile kullanıcı verisini saklamak ve güncellemek.
- JWT tabanlı kimlik doğrulamayı baştan sona uygulamak.

## Kullandığım Araçlar
- Go 1.21+
- Fiber v2
- MongoDB (lokal `mongodb://localhost:27017` veya Atlas)
- Postman (istekleri hızlı denemek için)

## Kurulum Adımlarım
1. Bağımlılıkları çektim:
   ```bash
   go mod download
   ```
2. `.env` dosyasını projenin köküne koyup bağlantı adresini tanımladım:
   ```text
   MONGODB_URI=mongodb://localhost:27017
   ```
3. MongoDB servisimin çalıştığından emin oldum (Compass veya `mongosh` ile kontrol edilebilir).

## Sunucuyu Çalıştırma
```bash
go run main.go, gor run .
```
Uygulama `http://localhost:3000` adresinde dinlemeye başlıyor.

## Postman ile Deneme Akışım
1. `POST http://localhost:3000/register`
   - Body: `raw` + `JSON`
   ```json
   { "email": "user@example.com", "password": "sifre123" }
   ```
   Bu istek yeni bir kullanıcıyı MongoDB'ye ekliyor.
2. `POST http://localhost:3000/login`
   - Aynı body ile giriş yapıyorum ve dönen JSON’daki `token` değerini alıyorum.
3. Token gerektiren isteklerde (örnek `GET http://localhost:3000/me`) Headers kısmına:
   - Key: `Authorization`
   - Value: `Bearer <token>`
4. Kitap ödünç alma (`POST http://localhost:3000/borrow`):
   ```json
   { "book": "Clean Code" }
   ```
   Token ile gönderince kullanıcıya kitap ekleniyor; 2 kitap sınırı var.
5. Kitap iade (`POST /return`) ve hesap silme (`DELETE /delete`) isteklerini de aynı token ile test ettim.

## Veritabanı Kontrolleri
- `library` veritabanı içinde `users` koleksiyonunda kullanıcıyı ve `books` alanındaki listeyi görebiliyorum.
- Örnek bir doküman şöyle görünüyor:
  ```json
  {
    "email": "eyluul@.gml.com",
    "password": "<bcrypt_hash>",
    "books": ["albert camus yabancı"]
  }
  ```

## JWT Üzerine Notum
`middleware.go` dosyasında `JwtSecret` adında bir sabit yerleştirdim. İlk denemelerde kod içinde kalması işimi kolaylaştırdı.

## Neler Öğrendim?
- Fiber ile route ve middleware yönetimi oldukça rahat.
- Bcrypt ile şifreleri hashlemek ve MongoDB'ye güvenli şekilde kaydetmek zor değil.
- JWT ile auth koyunca Postman testlerinde Authorization header’ını yönetmek kritik.

Bu proje benim için Go ile web servisleri öğrenmenin somut bir adımı oldu. 

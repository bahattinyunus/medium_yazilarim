# 🌟 Portfolio Generator - Otomatik Portfolio Oluşturucu

GitHub profilinizden otomatik olarak profesyonel bir portfolio websitesi oluşturan Node.js uygulaması.

## 🎯 Ne İşe Yarar?

- GitHub API'den kullanıcı bilgilerini çeker
- Son projelerinizi otomatik listeler
- Responsive HTML portfolio sayfası oluşturur
- Modern CSS styling ile güzel görünüm
- Tek komutla çalışır

## 🚀 Hızlı Başlangıç

### Gereksinimler
- Node.js (v12 veya üzeri)
- İnternet bağlantısı (GitHub API için)

### Kurulum
```bash
# Dosyaları indirin
git clone <repo-url>
cd portfolio-generator

# Çalıştırın
node portfolio-generator.js <github-kullanıcı-adı>
```

### Kullanım Örnekleri
```bash
# Kendi profiliniz için
node portfolio-generator.js bahattinyunus

# Ünlü geliştirici örnekleri
node portfolio-generator.js octocat
node portfolio-generator.js gaearon
node portfolio-generator.js sindresorhus
```

## 📁 Çıktı

Program `./portfolio/index.html` dosyası oluşturur:

```
portfolio/
└── index.html    # Tam portfolio websitesi
```

## 🎨 Özellikler

### ✅ Otomatik Veri Çekme
- Profil bilgileri (isim, bio, avatar)
- Repository statistics
- Son 6 proje (fork'lar hariç)
- Star/fork sayıları
- Kullanılan diller

### ✅ Modern Tasarım
- Responsive design (mobil uyumlu)
- Gradient hero section
- Card-based project layout
- Smooth hover animations
- Professional color scheme

### ✅ İletişim Bilgileri
- GitHub profil linki
- Email adresi (varsa)
- Kişisel website
- Lokasyon bilgisi

## 📋 Örnek Çıktı

```html
<!DOCTYPE html>
<html lang="tr">
<head>
    <title>Bahattin Yunus - Portfolio</title>
</head>
<body>
    <!-- Modern hero section -->
    <header class="hero">
        <img src="avatar.jpg" alt="Profile">
        <h1>Bahattin Yunus</h1>
        <p>Software Developer</p>
        <!-- Stats: Projeler, Takipçi, Takip -->
    </header>
    
    <!-- Project showcase -->
    <section class="projects">
        <!-- Son 6 proje card'ları -->
    </section>
    
    <!-- Contact info -->
    <section class="contact">
        <!-- GitHub, email, website, lokasyon -->
    </section>
</body>
</html>
```

## 🛠️ API Kullanımı

Program şu GitHub API endpoint'lerini kullanır:

```javascript
// Kullanıcı bilgileri
GET https://api.github.com/users/{username}

// Repository listesi
GET https://api.github.com/users/{username}/repos?sort=updated&per_page=6
```

### Rate Limiting
- GitHub API saatte 60 istek limiti
- Authenticated kullanım için token eklenebilir

## 🎯 Programatik Kullanım

```javascript
const PortfolioGenerator = require('./portfolio-generator');

// Generator oluştur
const generator = new PortfolioGenerator('username');

// Portfolio oluştur
generator.generatePortfolio()
    .then(() => console.log('Portfolio oluşturuldu!'))
    .catch(err => console.error('Hata:', err));
```

## 📚 Öğrenme Değeri

Bu proje şunları öğretir:

### Backend Konseptleri
- **HTTP Requests**: Native Node.js HTTPS modülü
- **API Integration**: GitHub REST API
- **Async/Await**: Promise-based programming
- **Error Handling**: Try-catch ve API hata yönetimi

### Frontend Konseptleri
- **Responsive Design**: CSS Grid ve Flexbox
- **CSS Animations**: Keyframes ve transitions
- **Template Strings**: Dynamic HTML generation
- **DOM Manipulation**: JavaScript interactivity

### Node.js Özellikleri
- **File System**: fs.promises ile dosya işlemleri
- **CLI Development**: process.argv ile komut satırı
- **Module System**: require/module.exports

## 🔧 Özelleştirme

### Farklı Template
```javascript
// Custom CSS stili
generator.getCSS = function() {
    return `
        /* Kendi CSS stiliniz */
        .hero { background: #your-color; }
    `;
};
```

### Daha Fazla Proje
```javascript
// 10 proje getir
const reposPromise = this.makeAPICall(
    `/users/${this.username}/repos?sort=updated&per_page=10`
);
```

### Authentication Token
```javascript
headers: {
    'User-Agent': 'Portfolio-Generator/1.0',
    'Authorization': `token ${process.env.GITHUB_TOKEN}`
}
```

## 🚀 Geliştirme Fikirleri

- [ ] **GitHub Pages Deploy**: Otomatik yayınlama
- [ ] **Theme System**: Farklı renk temaları
- [ ] **Analytics**: Google Analytics entegrasyonu
- [ ] **SEO**: Meta tags ve structured data
- [ ] **Multi-language**: İngilizce/Türkçe desteği
- [ ] **PDF Export**: Portfolio'yu PDF'e çevirme
- [ ] **Custom Sections**: Hakkında, Beceriler, Deneyim

## 🐛 Bilinen Sınırlamalar

- GitHub API rate limiting (60/saat)
- Sadece public repository'ler
- Basit hata yönetimi
- Tek sayfa çıktı
- Offline kullanım yok

## 🔍 Debug İpuçları

### API Hatası
```bash
# Kullanıcı adını kontrol edin
node portfolio-generator.js nonexistentuser
# ❌ GitHub veri çekme hatası: Not Found
```

### Network Hatası
```bash
# İnternet bağlantınızı kontrol edin
# API limitine takılmış olabilirsiniz
```

### Dosya Yazma Hatası
```bash
# Klasör izinlerini kontrol edin
chmod 755 ./
```

## 📖 İlgili Kaynaklar

- [GitHub API Docs](https://docs.github.com/en/rest)
- [Node.js HTTPS Module](https://nodejs.org/api/https.html)
- [CSS Grid Guide](https://css-tricks.com/snippets/css/complete-guide-grid/)
- [Responsive Web Design](https://web.dev/responsive-web-design-basics/)

Bu bir eğitim projesidir. Production kullanım için ek güvenlik ve optimizasyon gerekebilir.
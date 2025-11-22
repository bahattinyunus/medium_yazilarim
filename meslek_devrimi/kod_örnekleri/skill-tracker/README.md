# 🎯 Skill Tracker - Beceri Takip Sistemi

Bu basit Python uygulaması, yazılım geliştirici becerilerinizi takip etmenizi ve ilerlemenizi görselleştirmenizi sağlar.

## 📋 Ne İşe Yarar?

- Becerilerinizi 1-10 arası derecelendirme
- Kategorilere göre gruplama (Backend, Frontend, DevOps, vb.)
- İlerleme geçmişini kaydetme
- Basit raporlar oluşturma
- JSON formatında veri saklama

## 🚀 Nasıl Çalıştırılır?

```bash
# Dosyayı çalıştır
python skill_tracker.py

# Ya da executable olarak
chmod +x skill_tracker.py
./skill_tracker.py
```

## 📖 Kullanım Örnekleri

### Beceri Ekleme
```
Beceri adı: Python
Seviye (1-10): 7
Kategori: Backend
✅ Python becerisi seviye 7 olarak kaydedildi!
```

### Beceri Listesi Görünümü
```
🎯 Mevcut Beceriler:
--------------------------------------------------

📂 Backend:
   Python              [███████░░░] 7/10
   Node.js             [█████░░░░░] 5/10

📂 Frontend:
   React               [████████░░] 8/10
   JavaScript          [██████░░░░] 6/10
```

### Rapor Örneği
```
📊 Beceri Raporu
========================================
Toplam Beceri Sayısı: 4
Ortalama Seviye: 6.5/10

🏆 En Güçlü Beceriler:
   React: 8/10
   Python: 7/10
   JavaScript: 6/10

📈 Geliştirilmesi Gereken Alanlar:
   Node.js: 5/10
```

## 📁 Çıktı Dosyası

Program `skills.json` dosyasında verilerinizi saklar:

```json
{
  "skills": {
    "Python": {
      "level": 7,
      "category": "Backend",
      "last_updated": "2024-01-15T10:30:00",
      "created_date": "2024-01-10T09:15:00"
    }
  },
  "history": [
    {
      "skill": "Python",
      "level": 5,
      "category": "Backend",
      "timestamp": "2024-01-10T09:15:00",
      "action": "added"
    },
    {
      "skill": "Python",
      "level": 7,
      "category": "Backend", 
      "timestamp": "2024-01-15T10:30:00",
      "action": "updated"
    }
  ]
}
```

## 🛠️ Özelleştirme

### Farklı Veri Dosyası Kullanma
```python
tracker = SkillTracker("my_skills.json")
```

### Programatik Kullanım
```python
from skill_tracker import SkillTracker

# Tracker oluştur
tracker = SkillTracker()

# Beceri ekle
tracker.add_skill("Docker", 6, "DevOps")
tracker.add_skill("React", 8, "Frontend")

# Becerileri listele
tracker.list_skills()

# Rapor oluştur
tracker.generate_report()
```

## 📚 Öğrenme Değeri

Bu proje şu konuları öğretir:

- **File I/O**: JSON dosya okuma/yazma
- **Data Structures**: Dictionary ve list kullanımı
- **Error Handling**: Try-catch yapıları
- **CLI Development**: Kullanıcı etkileşimi
- **Data Visualization**: ASCII progress bar
- **Code Organization**: Class yapısı

## 🔧 Geliştirme Fikirleri

- Web interface ekle (Flask/FastAPI)
- Grafik raporlar (matplotlib)
- Beceri önerileri (AI entegrasyonu)
- Team skill tracking
- Export to PDF/Excel
- GitHub integration (commit analysis)

## 🐛 Bilinen Sınırlamalar

- Sadece lokal depolama (veritabanı yok)
- Basit validation
- Single user support
- ASCII-only progress bar

Bu bir eğitim amaçlı projedir. Production kullanım için ek güvenlik ve performans optimizasyonları gerekebilir.
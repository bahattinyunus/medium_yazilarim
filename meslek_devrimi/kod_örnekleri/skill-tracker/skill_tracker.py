#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Skill Tracker - Beceri Takip Sistemi
====================================

Bu basit araç, yazılım geliştirici becerilerini takip etmek ve
ilerlemeni görselleştirmek için tasarlanmıştır.

Özellikler:
- Beceri ekleme/güncelleme
- İlerleme takibi (1-10 arası)
- Zaman bazlı geçmiş
- Basit raporlama

Kullanım:
python skill_tracker.py
"""

import json
import os
from datetime import datetime
from typing import Dict, List, Optional


class SkillTracker:
    def __init__(self, data_file: str = "skills.json"):
        """
        Beceri takip sistemini başlat

        Args:
            data_file: Verilerin saklanacağı JSON dosyası
        """
        self.data_file = data_file
        self.skills = self.load_skills()

    def load_skills(self) -> Dict:
        """JSON dosyasından becerileri yükle"""
        if os.path.exists(self.data_file):
            try:
                with open(self.data_file, "r", encoding="utf-8") as f:
                    return json.load(f)
            except (json.JSONDecodeError, FileNotFoundError):
                print(f"⚠️  {self.data_file} okunamadı, yeni dosya oluşturuluyor...")

        return {"skills": {}, "history": []}

    def save_skills(self):
        """Becerileri JSON dosyasına kaydet"""
        try:
            with open(self.data_file, "w", encoding="utf-8") as f:
                json.dump(self.skills, f, indent=2, ensure_ascii=False)
            print("✅ Veriler kaydedildi!")
        except Exception as e:
            print(f"❌ Kayıt hatası: {e}")

    def add_skill(self, skill_name: str, level: int, category: str = "General"):
        """
        Yeni beceri ekle veya mevcut beceriyi güncelle

        Args:
            skill_name: Beceri adı (örn: "Python", "React")
            level: Beceri seviyesi (1-10 arası)
            category: Kategori (örn: "Backend", "Frontend", "Database")
        """
        if not 1 <= level <= 10:
            print("❌ Beceri seviyesi 1-10 arası olmalıdır!")
            return

        # Beceri bilgisini güncelle
        self.skills["skills"][skill_name] = {
            "level": level,
            "category": category,
            "last_updated": datetime.now().isoformat(),
            "created_date": self.skills["skills"]
            .get(skill_name, {})
            .get("created_date")
            or datetime.now().isoformat(),
        }

        # Geçmişe kaydet
        self.skills["history"].append(
            {
                "skill": skill_name,
                "level": level,
                "category": category,
                "timestamp": datetime.now().isoformat(),
                "action": "updated" if skill_name in self.skills["skills"] else "added",
            }
        )

        print(f"✅ {skill_name} becerisi seviye {level} olarak kaydedildi!")
        self.save_skills()

    def list_skills(self):
        """Tüm becerileri listele"""
        if not self.skills["skills"]:
            print("📝 Henüz hiç beceri eklenmemiş!")
            return

        print("\n🎯 Mevcut Beceriler:")
        print("-" * 50)

        # Kategoriye göre grupla
        categories = {}
        for skill, data in self.skills["skills"].items():
            category = data["category"]
            if category not in categories:
                categories[category] = []
            categories[category].append((skill, data))

        for category, skill_list in categories.items():
            print(f"\n📂 {category}:")
            for skill, data in sorted(
                skill_list, key=lambda x: x[1]["level"], reverse=True
            ):
                bar = "█" * data["level"] + "░" * (10 - data["level"])
                print(f"   {skill:<20} [{bar}] {data['level']}/10")

    def get_skill_progress(self, skill_name: str) -> List[Dict]:
        """Belirli bir becerinin ilerleme geçmişini getir"""
        return [
            entry for entry in self.skills["history"] if entry["skill"] == skill_name
        ]

    def generate_report(self):
        """Basit ilerleme raporu oluştur"""
        if not self.skills["skills"]:
            print("📝 Rapor için yeterli veri yok!")
            return

        print("\n📊 Beceri Raporu")
        print("=" * 40)

        total_skills = len(self.skills["skills"])
        avg_level = (
            sum(skill["level"] for skill in self.skills["skills"].values())
            / total_skills
        )

        print(f"Toplam Beceri Sayısı: {total_skills}")
        print(f"Ortalama Seviye: {avg_level:.1f}/10")

        # En yüksek beceriler
        top_skills = sorted(
            self.skills["skills"].items(), key=lambda x: x[1]["level"], reverse=True
        )[:3]

        print(f"\n🏆 En Güçlü Beceriler:")
        for skill, data in top_skills:
            print(f"   {skill}: {data['level']}/10")

        # Geliştirilmesi gereken alanlar
        weak_skills = [
            (skill, data)
            for skill, data in self.skills["skills"].items()
            if data["level"] < 5
        ]

        if weak_skills:
            print(f"\n📈 Geliştirilmesi Gereken Alanlar:")
            for skill, data in sorted(weak_skills, key=lambda x: x[1]["level"]):
                print(f"   {skill}: {data['level']}/10")


def main():
    """Ana menü sistemi"""
    tracker = SkillTracker()

    while True:
        print("\n" + "=" * 50)
        print("🎯 SKILL TRACKER - Beceri Takip Sistemi")
        print("=" * 50)
        print("1. Beceri Ekle/Güncelle")
        print("2. Becerileri Listele")
        print("3. Rapor Görüntüle")
        print("4. Çıkış")
        print("-" * 50)

        try:
            choice = input("Seçiminiz (1-4): ").strip()

            if choice == "1":
                skill_name = input("Beceri adı: ").strip()
                if not skill_name:
                    print("❌ Beceri adı boş olamaz!")
                    continue

                try:
                    level = int(input("Seviye (1-10): ").strip())
                except ValueError:
                    print("❌ Lütfen geçerli bir sayı girin!")
                    continue

                category = input("Kategori (opsiyonel): ").strip() or "General"
                tracker.add_skill(skill_name, level, category)

            elif choice == "2":
                tracker.list_skills()

            elif choice == "3":
                tracker.generate_report()

            elif choice == "4":
                print("👋 Görüşürüz!")
                break

            else:
                print("❌ Geçersiz seçim!")

        except KeyboardInterrupt:
            print("\n👋 Program sonlandırıldı!")
            break
        except Exception as e:
            print(f"❌ Beklenmeyen hata: {e}")


if __name__ == "__main__":
    main()

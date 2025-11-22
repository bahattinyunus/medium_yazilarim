#!/usr/bin/env node
/**
 * Portfolio Generator - Otomatik Portfolio Oluşturucu
 * ====================================================
 *
 * Bu Node.js uygulaması, GitHub profilinden bilgi çekerek
 * otomatik olarak HTML portfolio sayfası oluşturur.
 *
 * Özellikler:
 * - GitHub API entegrasyonu
 * - Responsive HTML template
 * - CSS styling
 * - Project showcase
 * - Contact form
 *
 * Kullanım:
 * node portfolio-generator.js <github-username>
 */

const fs = require('fs').promises;
const https = require('https');
const path = require('path');

class PortfolioGenerator {
    constructor(username) {
        this.username = username;
        this.githubAPI = 'https://api.github.com';
        this.outputDir = './portfolio';
    }

    /**
     * GitHub API'den veri çek
     */
    async fetchGitHubData() {
        try {
            console.log(`🔍 ${this.username} kullanıcısı araştırılıyor...`);

            const userPromise = this.makeAPICall(`/users/${this.username}`);
            const reposPromise = this.makeAPICall(`/users/${this.username}/repos?sort=updated&per_page=6`);

            const [userData, reposData] = await Promise.all([userPromise, reposPromise]);

            console.log(`✅ ${reposData.length} proje bulundu!`);

            return {
                user: userData,
                repos: reposData.filter(repo => !repo.fork) // Fork'ları hariç tut
            };
        } catch (error) {
            throw new Error(`GitHub veri çekme hatası: ${error.message}`);
        }
    }

    /**
     * GitHub API çağrısı yap
     */
    makeAPICall(endpoint) {
        return new Promise((resolve, reject) => {
            const options = {
                hostname: 'api.github.com',
                path: endpoint,
                headers: {
                    'User-Agent': 'Portfolio-Generator/1.0'
                }
            };

            const req = https.get(options, (res) => {
                let data = '';

                res.on('data', (chunk) => {
                    data += chunk;
                });

                res.on('end', () => {
                    try {
                        const jsonData = JSON.parse(data);

                        if (res.statusCode >= 400) {
                            reject(new Error(jsonData.message || 'API hatası'));
                        } else {
                            resolve(jsonData);
                        }
                    } catch (error) {
                        reject(new Error('JSON parse hatası'));
                    }
                });
            });

            req.on('error', (error) => {
                reject(error);
            });
        });
    }

    /**
     * HTML portfolio oluştur
     */
    generateHTML(data) {
        const { user, repos } = data;

        // Projeler için HTML
        const projectsHTML = repos.map(repo => `
            <div class="project-card">
                <h3>${repo.name}</h3>
                <p class="project-description">${repo.description || 'Açıklama mevcut değil'}</p>
                <div class="project-meta">
                    <span class="language">${repo.language || 'N/A'}</span>
                    <span class="stars">⭐ ${repo.stargazers_count}</span>
                    <span class="forks">🍴 ${repo.forks_count}</span>
                </div>
                <div class="project-links">
                    <a href="${repo.html_url}" target="_blank" class="btn-primary">Kodu Görüntüle</a>
                    ${repo.homepage ? `<a href="${repo.homepage}" target="_blank" class="btn-secondary">Demo</a>` : ''}
                </div>
            </div>
        `).join('');

        return `
<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>${user.name || user.login} - Portfolio</title>
    <style>
        ${this.getCSS()}
    </style>
</head>
<body>
    <!-- Header -->
    <header class="hero">
        <div class="container">
            <div class="hero-content">
                <img src="${user.avatar_url}" alt="Profile" class="profile-img">
                <h1>${user.name || user.login}</h1>
                <p class="tagline">${user.bio || 'Yazılım Geliştirici'}</p>
                <div class="hero-stats">
                    <div class="stat">
                        <span class="stat-number">${user.public_repos}</span>
                        <span class="stat-label">Proje</span>
                    </div>
                    <div class="stat">
                        <span class="stat-number">${user.followers}</span>
                        <span class="stat-label">Takipçi</span>
                    </div>
                    <div class="stat">
                        <span class="stat-number">${user.following}</span>
                        <span class="stat-label">Takip</span>
                    </div>
                </div>
            </div>
        </div>
    </header>

    <!-- Projects -->
    <section class="projects">
        <div class="container">
            <h2>🚀 Projelerim</h2>
            <div class="projects-grid">
                ${projectsHTML}
            </div>
            <div class="view-more">
                <a href="https://github.com/${user.login}" target="_blank" class="btn-outline">
                    Tüm Projeler GitHub'da
                </a>
            </div>
        </div>
    </section>

    <!-- Contact -->
    <section class="contact">
        <div class="container">
            <h2>📬 İletişim</h2>
            <div class="contact-info">
                <div class="contact-item">
                    <strong>GitHub:</strong>
                    <a href="https://github.com/${user.login}" target="_blank">@${user.login}</a>
                </div>
                ${user.email ? `
                <div class="contact-item">
                    <strong>Email:</strong>
                    <a href="mailto:${user.email}">${user.email}</a>
                </div>
                ` : ''}
                ${user.blog ? `
                <div class="contact-item">
                    <strong>Website:</strong>
                    <a href="${user.blog}" target="_blank">${user.blog}</a>
                </div>
                ` : ''}
                ${user.location ? `
                <div class="contact-item">
                    <strong>Konum:</strong> ${user.location}
                </div>
                ` : ''}
            </div>
        </div>
    </section>

    <!-- Footer -->
    <footer>
        <div class="container">
            <p>Bu portfolio <a href="https://github.com" target="_blank">Portfolio Generator</a> ile oluşturulmuştur.</p>
            <p>Son güncelleme: ${new Date().toLocaleDateString('tr-TR')}</p>
        </div>
    </footer>

    <script>
        // Basit animasyonlar
        document.addEventListener('DOMContentLoaded', function() {
            const cards = document.querySelectorAll('.project-card');
            cards.forEach((card, index) => {
                card.style.animationDelay = (index * 0.1) + 's';
                card.classList.add('fade-in');
            });
        });
    </script>
</body>
</html>`;
    }

    /**
     * CSS stilleri
     */
    getCSS() {
        return `
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            line-height: 1.6;
            color: #333;
            background-color: #f8fafc;
        }

        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 0 20px;
        }

        .hero {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 80px 0;
            text-align: center;
        }

        .profile-img {
            width: 120px;
            height: 120px;
            border-radius: 50%;
            border: 4px solid rgba(255,255,255,0.3);
            margin-bottom: 20px;
        }

        .hero h1 {
            font-size: 2.5em;
            margin-bottom: 10px;
        }

        .tagline {
            font-size: 1.2em;
            opacity: 0.9;
            margin-bottom: 30px;
        }

        .hero-stats {
            display: flex;
            justify-content: center;
            gap: 30px;
            margin-top: 30px;
        }

        .stat {
            text-align: center;
        }

        .stat-number {
            display: block;
            font-size: 2em;
            font-weight: bold;
        }

        .stat-label {
            font-size: 0.9em;
            opacity: 0.8;
        }

        .projects {
            padding: 80px 0;
        }

        .projects h2 {
            text-align: center;
            font-size: 2.2em;
            margin-bottom: 50px;
            color: #2d3748;
        }

        .projects-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
            gap: 30px;
            margin-bottom: 50px;
        }

        .project-card {
            background: white;
            padding: 30px;
            border-radius: 12px;
            box-shadow: 0 4px 6px rgba(0,0,0,0.1);
            transition: transform 0.3s ease, box-shadow 0.3s ease;
            opacity: 0;
            transform: translateY(20px);
        }

        .project-card:hover {
            transform: translateY(-5px);
            box-shadow: 0 8px 25px rgba(0,0,0,0.15);
        }

        .project-card.fade-in {
            animation: fadeInUp 0.6s ease forwards;
        }

        @keyframes fadeInUp {
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }

        .project-card h3 {
            color: #2d3748;
            margin-bottom: 15px;
            font-size: 1.3em;
        }

        .project-description {
            color: #4a5568;
            margin-bottom: 20px;
            line-height: 1.5;
        }

        .project-meta {
            display: flex;
            gap: 15px;
            margin-bottom: 20px;
            font-size: 0.9em;
            color: #718096;
        }

        .language {
            background: #edf2f7;
            padding: 4px 8px;
            border-radius: 4px;
            font-weight: 500;
        }

        .project-links {
            display: flex;
            gap: 10px;
        }

        .btn-primary, .btn-secondary, .btn-outline {
            padding: 8px 16px;
            border-radius: 6px;
            text-decoration: none;
            font-weight: 500;
            font-size: 0.9em;
            transition: all 0.2s ease;
            display: inline-block;
        }

        .btn-primary {
            background: #667eea;
            color: white;
        }

        .btn-primary:hover {
            background: #5a6fd8;
        }

        .btn-secondary {
            background: #48bb78;
            color: white;
        }

        .btn-secondary:hover {
            background: #38a169;
        }

        .btn-outline {
            border: 2px solid #667eea;
            color: #667eea;
            background: transparent;
        }

        .btn-outline:hover {
            background: #667eea;
            color: white;
        }

        .view-more {
            text-align: center;
        }

        .contact {
            background: white;
            padding: 80px 0;
        }

        .contact h2 {
            text-align: center;
            font-size: 2.2em;
            margin-bottom: 50px;
            color: #2d3748;
        }

        .contact-info {
            max-width: 600px;
            margin: 0 auto;
        }

        .contact-item {
            padding: 15px 0;
            border-bottom: 1px solid #e2e8f0;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }

        .contact-item:last-child {
            border-bottom: none;
        }

        .contact-item a {
            color: #667eea;
            text-decoration: none;
        }

        .contact-item a:hover {
            text-decoration: underline;
        }

        footer {
            background: #2d3748;
            color: white;
            text-align: center;
            padding: 40px 0;
        }

        footer a {
            color: #667eea;
            text-decoration: none;
        }

        footer p {
            margin-bottom: 10px;
        }

        @media (max-width: 768px) {
            .hero-stats {
                flex-direction: column;
                gap: 15px;
            }

            .projects-grid {
                grid-template-columns: 1fr;
            }

            .contact-item {
                flex-direction: column;
                align-items: flex-start;
                gap: 5px;
            }
        }
        `;
    }

    /**
     * Çıktı klasörü oluştur ve portfolio dosyalarını yaz
     */
    async generatePortfolio() {
        try {
            // GitHub verilerini çek
            const data = await this.fetchGitHubData();

            // HTML oluştur
            const htmlContent = this.generateHTML(data);

            // Çıktı klasörü oluştur
            await fs.mkdir(this.outputDir, { recursive: true });

            // HTML dosyasını yaz
            const htmlPath = path.join(this.outputDir, 'index.html');
            await fs.writeFile(htmlPath, htmlContent, 'utf8');

            console.log(`✅ Portfolio oluşturuldu: ${htmlPath}`);
            console.log(`🌐 Tarayıcıda açmak için: file://${path.resolve(htmlPath)}`);

        } catch (error) {
            console.error(`❌ Hata: ${error.message}`);
            process.exit(1);
        }
    }
}

// CLI kullanımı
function main() {
    const args = process.argv.slice(2);

    if (args.length === 0) {
        console.log(`
🎯 Portfolio Generator - Kullanım:

node portfolio-generator.js <github-username>

Örnekler:
  node portfolio-generator.js octocat
  node portfolio-generator.js torvalds
  node portfolio-generator.js gaearon

Bu araç GitHub API kullanarak otomatik portfolio oluşturur.
        `);
        process.exit(1);
    }

    const username = args[0];
    const generator = new PortfolioGenerator(username);

    console.log(`🚀 ${username} için portfolio oluşturuluyor...`);
    generator.generatePortfolio();
}

// Direkt çalıştırıldığında main fonksiyonunu çağır
if (require.main === module) {
    main();
}

module.exports = PortfolioGenerator;

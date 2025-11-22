package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// JobListing represents a job posting
type JobListing struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Company     string   `json:"company"`
	Location    string   `json:"location"`
	Description string   `json:"description"`
	Requirements []string `json:"requirements"`
	Skills      []string `json:"skills"`
	SalaryMin   int      `json:"salary_min"`
	SalaryMax   int      `json:"salary_max"`
	PostedDate  string   `json:"posted_date"`
	Remote      bool     `json:"remote"`
}

// JobAnalyzer analyzes job market data
type JobAnalyzer struct {
	Jobs []JobListing `json:"jobs"`
}

// SkillStats represents skill statistics
type SkillStats struct {
	Skill string `json:"skill"`
	Count int    `json:"count"`
	Percentage float64 `json:"percentage"`
}

// SalaryStats represents salary statistics
type SalaryStats struct {
	MinSalary    int     `json:"min_salary"`
	MaxSalary    int     `json:"max_salary"`
	AverageSalary float64 `json:"average_salary"`
	MedianSalary int     `json:"median_salary"`
}

// CompanyStats represents company statistics
type CompanyStats struct {
	Company string `json:"company"`
	JobCount int   `json:"job_count"`
}

// NewJobAnalyzer creates a new job analyzer instance
func NewJobAnalyzer() *JobAnalyzer {
	return &JobAnalyzer{
		Jobs: make([]JobListing, 0),
	}
}

// LoadSampleData loads sample job data for demonstration
func (ja *JobAnalyzer) LoadSampleData() {
	sampleJobs := []JobListing{
		{
			ID: "1",
			Title: "Senior Software Engineer",
			Company: "TechCorp",
			Location: "Istanbul, Turkey",
			Description: "We are looking for a senior software engineer with experience in Go, Python, and cloud technologies.",
			Skills: []string{"Go", "Python", "Docker", "Kubernetes", "AWS", "PostgreSQL"},
			SalaryMin: 15000,
			SalaryMax: 25000,
			PostedDate: "2024-01-15",
			Remote: true,
		},
		{
			ID: "2",
			Title: "Frontend Developer",
			Company: "StartupXYZ",
			Location: "Ankara, Turkey",
			Description: "Join our team as a frontend developer. React and TypeScript experience required.",
			Skills: []string{"React", "TypeScript", "JavaScript", "CSS", "HTML", "Redux"},
			SalaryMin: 10000,
			SalaryMax: 18000,
			PostedDate: "2024-01-14",
			Remote: false,
		},
		{
			ID: "3",
			Title: "DevOps Engineer",
			Company: "CloudSolutions",
			Location: "Izmir, Turkey",
			Description: "Looking for DevOps engineer to manage our cloud infrastructure and CI/CD pipelines.",
			Skills: []string{"Docker", "Kubernetes", "AWS", "Terraform", "Jenkins", "Linux", "Python"},
			SalaryMin: 18000,
			SalaryMax: 30000,
			PostedDate: "2024-01-13",
			Remote: true,
		},
		{
			ID: "4",
			Title: "Full Stack Developer",
			Company: "WebAgency",
			Location: "Bursa, Turkey",
			Description: "Full stack developer position with Node.js backend and React frontend experience.",
			Skills: []string{"Node.js", "React", "MongoDB", "Express", "JavaScript", "TypeScript"},
			SalaryMin: 12000,
			SalaryMax: 20000,
			PostedDate: "2024-01-12",
			Remote: true,
		},
		{
			ID: "5",
			Title: "Data Scientist",
			Company: "DataCorp",
			Location: "Istanbul, Turkey",
			Description: "Data scientist role focusing on machine learning and data analysis.",
			Skills: []string{"Python", "R", "TensorFlow", "Pandas", "NumPy", "SQL", "Machine Learning"},
			SalaryMin: 16000,
			SalaryMax: 28000,
			PostedDate: "2024-01-11",
			Remote: false,
		},
		{
			ID: "6",
			Title: "Backend Developer",
			Company: "ApiFirst",
			Location: "Antalya, Turkey",
			Description: "Backend developer specializing in API development and microservices.",
			Skills: []string{"Java", "Spring Boot", "PostgreSQL", "Redis", "Docker", "Microservices"},
			SalaryMin: 13000,
			SalaryMax: 22000,
			PostedDate: "2024-01-10",
			Remote: true,
		},
	}

	ja.Jobs = sampleJobs
}

// LoadFromFile loads job data from JSON file
func (ja *JobAnalyzer) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("dosya açılamadı: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("dosya okunamadı: %v", err)
	}

	err = json.Unmarshal(data, &ja.Jobs)
	if err != nil {
		return fmt.Errorf("JSON parse hatası: %v", err)
	}

	return nil
}

// SaveToFile saves job data to JSON file
func (ja *JobAnalyzer) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(ja.Jobs, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON serialize hatası: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("dosya yazma hatası: %v", err)
	}

	return nil
}

// AnalyzeSkills analyzes the most demanded skills
func (ja *JobAnalyzer) AnalyzeSkills() []SkillStats {
	skillCount := make(map[string]int)
	totalJobs := len(ja.Jobs)

	// Count skills
	for _, job := range ja.Jobs {
		for _, skill := range job.Skills {
			skillCount[strings.ToLower(skill)]++
		}
	}

	// Convert to slice and sort
	var skills []SkillStats
	for skill, count := range skillCount {
		percentage := float64(count) / float64(totalJobs) * 100
		skills = append(skills, SkillStats{
			Skill:      skill,
			Count:      count,
			Percentage: percentage,
		})
	}

	// Sort by count (descending)
	sort.Slice(skills, func(i, j int) bool {
		return skills[i].Count > skills[j].Count
	})

	return skills
}

// AnalyzeSalaries analyzes salary statistics
func (ja *JobAnalyzer) AnalyzeSalaries() SalaryStats {
	if len(ja.Jobs) == 0 {
		return SalaryStats{}
	}

	var salaries []int
	minSal := ja.Jobs[0].SalaryMin
	maxSal := ja.Jobs[0].SalaryMax
	totalSalary := 0

	for _, job := range ja.Jobs {
		if job.SalaryMin > 0 {
			salaries = append(salaries, job.SalaryMin)
			totalSalary += job.SalaryMin
		}
		if job.SalaryMax > 0 {
			salaries = append(salaries, job.SalaryMax)
			totalSalary += job.SalaryMax
		}

		if job.SalaryMin > 0 && job.SalaryMin < minSal {
			minSal = job.SalaryMin
		}
		if job.SalaryMax > maxSal {
			maxSal = job.SalaryMax
		}
	}

	// Calculate average
	averageSalary := float64(totalSalary) / float64(len(salaries))

	// Calculate median
	sort.Ints(salaries)
	medianSalary := salaries[len(salaries)/2]

	return SalaryStats{
		MinSalary:     minSal,
		MaxSalary:     maxSal,
		AverageSalary: averageSalary,
		MedianSalary:  medianSalary,
	}
}

// AnalyzeCompanies analyzes which companies post the most jobs
func (ja *JobAnalyzer) AnalyzeCompanies() []CompanyStats {
	companyCount := make(map[string]int)

	for _, job := range ja.Jobs {
		companyCount[job.Company]++
	}

	var companies []CompanyStats
	for company, count := range companyCount {
		companies = append(companies, CompanyStats{
			Company:  company,
			JobCount: count,
		})
	}

	// Sort by job count (descending)
	sort.Slice(companies, func(i, j int) bool {
		return companies[i].JobCount > companies[j].JobCount
	})

	return companies
}

// FilterBySkill filters jobs that require a specific skill
func (ja *JobAnalyzer) FilterBySkill(skill string) []JobListing {
	var filteredJobs []JobListing
	skill = strings.ToLower(skill)

	for _, job := range ja.Jobs {
		for _, jobSkill := range job.Skills {
			if strings.ToLower(jobSkill) == skill {
				filteredJobs = append(filteredJobs, job)
				break
			}
		}
	}

	return filteredJobs
}

// FilterBySalaryRange filters jobs by salary range
func (ja *JobAnalyzer) FilterBySalaryRange(minSalary, maxSalary int) []JobListing {
	var filteredJobs []JobListing

	for _, job := range ja.Jobs {
		if job.SalaryMin >= minSalary && job.SalaryMax <= maxSalary {
			filteredJobs = append(filteredJobs, job)
		}
	}

	return filteredJobs
}

// FilterByLocation filters jobs by location
func (ja *JobAnalyzer) FilterByLocation(location string) []JobListing {
	var filteredJobs []JobListing
	location = strings.ToLower(location)

	for _, job := range ja.Jobs {
		if strings.Contains(strings.ToLower(job.Location), location) {
			filteredJobs = append(filteredJobs, job)
		}
	}

	return filteredJobs
}

// FilterByRemote filters remote jobs
func (ja *JobAnalyzer) FilterByRemote() []JobListing {
	var filteredJobs []JobListing

	for _, job := range ja.Jobs {
		if job.Remote {
			filteredJobs = append(filteredJobs, job)
		}
	}

	return filteredJobs
}

// GenerateReport generates a comprehensive analysis report
func (ja *JobAnalyzer) GenerateReport() {
	fmt.Println("🔍 İŞ PAZARI ANALİZ RAPORU")
	fmt.Println("=" + strings.Repeat("=", 50))

	fmt.Printf("📊 Toplam İş İlanı: %d\n", len(ja.Jobs))
	fmt.Printf("📅 Analiz Tarihi: %s\n\n", time.Now().Format("2006-01-02 15:04"))

	// Skills Analysis
	fmt.Println("🎯 EN POPÜLER YETENEKLER")
	fmt.Println("-" + strings.Repeat("-", 30))
	skills := ja.AnalyzeSkills()
	for i, skill := range skills[:min(10, len(skills))] {
		fmt.Printf("%2d. %-15s %3d iş (%.1f%%)\n",
			i+1, strings.Title(skill.Skill), skill.Count, skill.Percentage)
	}

	// Salary Analysis
	fmt.Println("\n💰 MAAŞ ANALİZİ")
	fmt.Println("-" + strings.Repeat("-", 20))
	salaryStats := ja.AnalyzeSalaries()
	fmt.Printf("En Düşük: %,d TL\n", salaryStats.MinSalary)
	fmt.Printf("En Yüksek: %,d TL\n", salaryStats.MaxSalary)
	fmt.Printf("Ortalama:  %,.0f TL\n", salaryStats.AverageSalary)
	fmt.Printf("Medyan:    %,d TL\n", salaryStats.MedianSalary)

	// Company Analysis
	fmt.Println("\n🏢 EN AKTİF ŞİRKETLER")
	fmt.Println("-" + strings.Repeat("-", 25))
	companies := ja.AnalyzeCompanies()
	for i, company := range companies[:min(5, len(companies))] {
		fmt.Printf("%d. %s (%d iş)\n", i+1, company.Company, company.JobCount)
	}

	// Remote Work Analysis
	remoteJobs := ja.FilterByRemote()
	remotePercentage := float64(len(remoteJobs)) / float64(len(ja.Jobs)) * 100
	fmt.Printf("\n🏠 Uzaktan Çalışma: %d iş (%.1f%%)\n", len(remoteJobs), remotePercentage)

	// Location Analysis
	fmt.Println("\n📍 LOKASYON DAĞILIMI")
	fmt.Println("-" + strings.Repeat("-", 20))
	locationCount := make(map[string]int)
	for _, job := range ja.Jobs {
		// Extract city from location
		parts := strings.Split(job.Location, ",")
		if len(parts) > 0 {
			city := strings.TrimSpace(parts[0])
			locationCount[city]++
		}
	}

	type locationStat struct {
		Location string
		Count    int
	}

	var locations []locationStat
	for loc, count := range locationCount {
		locations = append(locations, locationStat{loc, count})
	}

	sort.Slice(locations, func(i, j int) bool {
		return locations[i].Count > locations[j].Count
	})

	for i, loc := range locations[:min(5, len(locations))] {
		fmt.Printf("%d. %s (%d iş)\n", i+1, loc.Location, loc.Count)
	}
}

// SearchJobs searches jobs by keyword in title or description
func (ja *JobAnalyzer) SearchJobs(keyword string) []JobListing {
	var results []JobListing
	keyword = strings.ToLower(keyword)

	for _, job := range ja.Jobs {
		titleMatch := strings.Contains(strings.ToLower(job.Title), keyword)
		descMatch := strings.Contains(strings.ToLower(job.Description), keyword)

		if titleMatch || descMatch {
			results = append(results, job)
		}
	}

	return results
}

// DisplayJobs displays jobs in a formatted way
func (ja *JobAnalyzer) DisplayJobs(jobs []JobListing) {
	if len(jobs) == 0 {
		fmt.Println("❌ Hiç iş bulunamadı!")
		return
	}

	fmt.Printf("📋 %d İş İlanı Bulundu:\n\n", len(jobs))

	for i, job := range jobs {
		fmt.Printf("🔹 %d. %s\n", i+1, job.Title)
		fmt.Printf("   🏢 %s | 📍 %s\n", job.Company, job.Location)
		fmt.Printf("   💰 %,d - %,d TL", job.SalaryMin, job.SalaryMax)
		if job.Remote {
			fmt.Print(" | 🏠 Uzaktan")
		}
		fmt.Println()
		fmt.Printf("   🛠️  %s\n", strings.Join(job.Skills, ", "))
		fmt.Println()
	}
}

// min helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Interactive menu system
func (ja *JobAnalyzer) RunInteractive() {
	for {
		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Println("🔍 İŞ PAZARI ANALİZCİSİ")
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("1. 📊 Genel Rapor")
		fmt.Println("2. 🎯 Yetenek Bazlı Arama")
		fmt.Println("3. 💰 Maaş Aralığı Filtresi")
		fmt.Println("4. 📍 Lokasyon Filtresi")
		fmt.Println("5. 🏠 Uzaktan Çalışma İşleri")
		fmt.Println("6. 🔎 Anahtar Kelime Arama")
		fmt.Println("7. 📁 Veri Yönetimi")
		fmt.Println("8. ❌ Çıkış")
		fmt.Println(strings.Repeat("-", 50))

		fmt.Print("Seçiminizi yapın (1-8): ")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			ja.GenerateReport()
		case "2":
			fmt.Print("Aranacak yetenek: ")
			var skill string
			fmt.Scanln(&skill)
			jobs := ja.FilterBySkill(skill)
			ja.DisplayJobs(jobs)
		case "3":
			fmt.Print("Minimum maaş: ")
			var minStr string
			fmt.Scanln(&minStr)
			fmt.Print("Maksimum maaş: ")
			var maxStr string
			fmt.Scanln(&maxStr)

			min, err1 := strconv.Atoi(minStr)
			max, err2 := strconv.Atoi(maxStr)

			if err1 != nil || err2 != nil {
				fmt.Println("❌ Geçersiz maaş değeri!")
				continue
			}

			jobs := ja.FilterBySalaryRange(min, max)
			ja.DisplayJobs(jobs)
		case "4":
			fmt.Print("Lokasyon: ")
			var location string
			fmt.Scanln(&location)
			jobs := ja.FilterByLocation(location)
			ja.DisplayJobs(jobs)
		case "5":
			jobs := ja.FilterByRemote()
			ja.DisplayJobs(jobs)
		case "6":
			fmt.Print("Anahtar kelime: ")
			var keyword string
			fmt.Scanln(&keyword)
			jobs := ja.SearchJobs(keyword)
			ja.DisplayJobs(jobs)
		case "7":
			ja.dataManagementMenu()
		case "8":
			fmt.Println("👋 Güle güle!")
			return
		default:
			fmt.Println("❌ Geçersiz seçim!")
		}

		fmt.Print("\nDevam etmek için Enter'a basın...")
		fmt.Scanln()
	}
}

func (ja *JobAnalyzer) dataManagementMenu() {
	fmt.Println("\n📁 VERİ YÖNETİMİ")
	fmt.Println("1. 📥 Dosyadan Yükle")
	fmt.Println("2. 💾 Dosyaya Kaydet")
	fmt.Println("3. 🔄 Örnek Veri Yükle")
	fmt.Print("Seçim: ")

	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		fmt.Print("Dosya adı: ")
		var filename string
		fmt.Scanln(&filename)
		err := ja.LoadFromFile(filename)
		if err != nil {
			fmt.Printf("❌ Hata: %v\n", err)
		} else {
			fmt.Printf("✅ %d iş yüklendi!\n", len(ja.Jobs))
		}
	case "2":
		fmt.Print("Dosya adı: ")
		var filename string
		fmt.Scanln(&filename)
		err := ja.SaveToFile(filename)
		if err != nil {
			fmt.Printf("❌ Hata: %v\n", err)
		} else {
			fmt.Println("✅ Veriler kaydedildi!")
		}
	case "3":
		ja.LoadSampleData()
		fmt.Printf("✅ %d örnek iş yüklendi!\n", len(ja.Jobs))
	}
}

func main() {
	fmt.Println("🚀 Job Analyzer - İş Piyasası Analiz Aracı")
	fmt.Println("==========================================")

	analyzer := NewJobAnalyzer()

	// Load sample data by default
	analyzer.LoadSampleData()
	fmt.Printf("📊 %d örnek iş verisi yüklendi.\n", len(analyzer.Jobs))

	// Run interactive mode
	analyzer.RunInteractive()
}

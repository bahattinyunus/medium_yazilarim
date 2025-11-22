#!/bin/bash

# Smart Commit - Akıllı Git Commit Aracı
# =====================================
#
# Bu script, git commit sürecini otomatikleştirir ve standardize eder.
# Conventional commits standardını destekler ve otomatik validasyon yapar.
#
# Kullanım:
#   ./smart-commit.sh "feat: user authentication system"
#   ./smart-commit.sh --interactive
#   ./smart-commit.sh --quick

set -e  # Exit on error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
VALID_TYPES=("feat" "fix" "docs" "style" "refactor" "perf" "test" "chore" "ci" "build" "revert")
MIN_MESSAGE_LENGTH=10
MAX_MESSAGE_LENGTH=100

# Functions
print_usage() {
    echo -e "${BLUE}Smart Commit - Akıllı Git Commit Aracı${NC}"
    echo
    echo "Kullanım:"
    echo "  $0 \"commit-message\"       # Direkt commit"
    echo "  $0 --interactive           # Etkileşimli mod"
    echo "  $0 --quick                 # Hızlı commit (otomatik message)"
    echo "  $0 --help                  # Bu yardım mesajı"
    echo
    echo "Conventional Commit Formatı:"
    echo "  type(scope): description"
    echo
    echo "Geçerli tipler:"
    printf "  "
    for type in "${VALID_TYPES[@]}"; do
        echo -n "$type "
    done
    echo
    echo
    echo "Örnekler:"
    echo "  $0 \"feat(auth): add login functionality\""
    echo "  $0 \"fix: resolve memory leak in parser\""
    echo "  $0 \"docs: update API documentation\""
}

print_error() {
    echo -e "${RED}❌ Hata: $1${NC}" >&2
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

# Check if we're in a git repository
check_git_repo() {
    if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
        print_error "Bu dizin bir Git repository değil!"
        exit 1
    fi
}

# Check for uncommitted changes
check_changes() {
    if git diff-index --quiet HEAD -- 2>/dev/null; then
        if git diff --cached --quiet; then
            print_warning "Commit edilecek değişiklik bulunamadı!"

            # Show untracked files
            if [ -n "$(git ls-files --others --exclude-standard)" ]; then
                echo
                print_info "Tracked olmayan dosyalar:"
                git ls-files --others --exclude-standard | sed 's/^/  /'
                echo
                read -p "Bu dosyaları eklemek ister misiniz? (y/N): " -n 1 -r
                echo
                if [[ $REPLY =~ ^[Yy]$ ]]; then
                    git add .
                    print_success "Tüm dosyalar eklendi"
                else
                    exit 1
                fi
            else
                exit 1
            fi
        fi
    fi
}

# Validate commit message format
validate_message() {
    local message="$1"

    # Check length
    if [ ${#message} -lt $MIN_MESSAGE_LENGTH ]; then
        print_error "Commit mesajı çok kısa! (minimum $MIN_MESSAGE_LENGTH karakter)"
        return 1
    fi

    if [ ${#message} -gt $MAX_MESSAGE_LENGTH ]; then
        print_error "Commit mesajı çok uzun! (maksimum $MAX_MESSAGE_LENGTH karakter)"
        return 1
    fi

    # Check conventional commit format
    if [[ $message =~ ^([a-z]+)(\(.+\))?: .+ ]]; then
        local type="${BASH_REMATCH[1]}"

        # Check if type is valid
        local valid_type=false
        for valid in "${VALID_TYPES[@]}"; do
            if [ "$type" = "$valid" ]; then
                valid_type=true
                break
            fi
        done

        if [ "$valid_type" = false ]; then
            print_error "Geçersiz commit tipi: '$type'"
            echo "Geçerli tipler: ${VALID_TYPES[*]}"
            return 1
        fi

        print_success "Commit mesajı format kontrolü başarılı"
        return 0
    else
        print_warning "Conventional commit formatı kullanılmıyor"
        read -p "Yine de devam etmek istiyor musunuz? (y/N): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            return 0
        else
            return 1
        fi
    fi
}

# Generate automatic commit message based on changes
generate_auto_message() {
    local changes=$(git diff --cached --name-only | wc -l)
    local added=$(git diff --cached --diff-filter=A --name-only | wc -l)
    local modified=$(git diff --cached --diff-filter=M --name-only | wc -l)
    local deleted=$(git diff --cached --diff-filter=D --name-only | wc -l)

    local message=""

    if [ $added -gt 0 ] && [ $modified -eq 0 ] && [ $deleted -eq 0 ]; then
        if [ $added -eq 1 ]; then
            local filename=$(git diff --cached --diff-filter=A --name-only)
            message="feat: add $filename"
        else
            message="feat: add $added new files"
        fi
    elif [ $added -eq 0 ] && [ $modified -gt 0 ] && [ $deleted -eq 0 ]; then
        if [ $modified -eq 1 ]; then
            local filename=$(git diff --cached --diff-filter=M --name-only)
            message="fix: update $filename"
        else
            message="fix: update $modified files"
        fi
    elif [ $added -eq 0 ] && [ $modified -eq 0 ] && [ $deleted -gt 0 ]; then
        if [ $deleted -eq 1 ]; then
            local filename=$(git diff --cached --diff-filter=D --name-only)
            message="chore: remove $filename"
        else
            message="chore: remove $deleted files"
        fi
    else
        message="chore: update $changes files"
    fi

    echo "$message"
}

# Show commit preview
show_preview() {
    echo
    echo -e "${BLUE}📋 Commit Önizleme${NC}"
    echo "==================="

    echo
    echo -e "${YELLOW}Değişiklikler:${NC}"
    git status --porcelain --cached | while read line; do
        local status="${line:0:2}"
        local file="${line:3}"
        case "$status" in
            "A ") echo -e "  ${GREEN}+${NC} $file" ;;
            "M ") echo -e "  ${YELLOW}~${NC} $file" ;;
            "D ") echo -e "  ${RED}-${NC} $file" ;;
            *) echo -e "  ${BLUE}?${NC} $file" ;;
        esac
    done

    echo
    echo -e "${YELLOW}İstatistikler:${NC}"
    git diff --cached --stat
    echo
}

# Interactive mode
interactive_mode() {
    echo -e "${BLUE}🤖 Etkileşimli Commit Modu${NC}"
    echo "=========================="

    # Show current status
    show_preview

    # Get commit type
    echo
    echo "Commit tipi seçin:"
    for i in "${!VALID_TYPES[@]}"; do
        printf "%2d) %s\n" $((i+1)) "${VALID_TYPES[$i]}"
    done

    while true; do
        read -p "Seçim (1-${#VALID_TYPES[@]}): " -r type_choice
        if [[ $type_choice =~ ^[0-9]+$ ]] && [ $type_choice -ge 1 ] && [ $type_choice -le ${#VALID_TYPES[@]} ]; then
            selected_type="${VALID_TYPES[$((type_choice-1))]}"
            break
        else
            print_error "Geçersiz seçim!"
        fi
    done

    # Get scope (optional)
    echo
    read -p "Scope (opsiyonel, örn: auth, api, ui): " -r scope

    # Get description
    echo
    read -p "Açıklama: " -r description

    # Build commit message
    if [ -n "$scope" ]; then
        commit_message="$selected_type($scope): $description"
    else
        commit_message="$selected_type: $description"
    fi

    echo
    echo -e "${YELLOW}Commit mesajı: ${NC}$commit_message"

    read -p "Onaylıyor musunuz? (Y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Nn]$ ]]; then
        print_warning "Commit iptal edildi"
        exit 1
    fi

    echo "$commit_message"
}

# Perform the actual commit
do_commit() {
    local message="$1"
    local show_preview="$2"

    if [ "$show_preview" = "true" ]; then
        show_preview
        echo -e "${YELLOW}Commit mesajı: ${NC}$message"
        echo
        read -p "Commit yapmak istediğinizden emin misiniz? (Y/n): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Nn]$ ]]; then
            print_warning "Commit iptal edildi"
            exit 1
        fi
    fi

    # Perform commit
    if git commit -m "$message"; then
        print_success "Commit başarılı!"

        # Show commit info
        echo
        git log -1 --oneline

        # Ask about push
        if git remote >/dev/null 2>&1; then
            echo
            read -p "Değişiklikleri push etmek ister misiniz? (y/N): " -n 1 -r
            echo
            if [[ $REPLY =~ ^[Yy]$ ]]; then
                current_branch=$(git branch --show-current)
                if git push origin "$current_branch"; then
                    print_success "Push başarılı!"
                else
                    print_error "Push başarısız!"
                fi
            fi
        fi
    else
        print_error "Commit başarısız!"
        exit 1
    fi
}

# Main execution
main() {
    check_git_repo
    check_changes

    case "${1:-}" in
        "--help"|"-h")
            print_usage
            exit 0
            ;;
        "--interactive"|"-i")
            commit_message=$(interactive_mode)
            if validate_message "$commit_message"; then
                do_commit "$commit_message" false
            fi
            ;;
        "--quick"|"-q")
            commit_message=$(generate_auto_message)
            print_info "Otomatik mesaj: $commit_message"
            if validate_message "$commit_message"; then
                do_commit "$commit_message" true
            fi
            ;;
        "")
            print_error "Commit mesajı gerekli!"
            echo
            print_usage
            exit 1
            ;;
        *)
            commit_message="$1"
            if validate_message "$commit_message"; then
                do_commit "$commit_message" true
            fi
            ;;
    esac
}

# Run main function with all arguments
main "$@"

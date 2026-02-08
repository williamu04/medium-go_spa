const API_BASE = 'http://localhost:8080/v1/api';
let currentUser = null;
let authToken = null;

// Initialize
document.addEventListener('DOMContentLoaded', () => {
    checkAuth();
    navigateTo('home');
});

// Auth Functions
function checkAuth() {
    authToken = localStorage.getItem('authToken');
    const userData = localStorage.getItem('userData');
    
    if (authToken && userData) {
        currentUser = JSON.parse(userData);
        updateAuthUI();
    }
}

function updateAuthUI() {
    const authBtn = document.getElementById('authBtn');
    const profileBtn = document.getElementById('profileBtn');
    
    if (currentUser) {
        authBtn.textContent = 'Keluar';
        authBtn.onclick = logout;
        profileBtn.classList.remove('hidden');
    } else {
        authBtn.textContent = 'Masuk';
        authBtn.onclick = () => navigateTo('login');
        profileBtn.classList.add('hidden');
    }
}

function handleAuth() {
    if (currentUser) {
        logout();
    } else {
        navigateTo('login');
    }
}

function logout() {
    localStorage.removeItem('authToken');
    localStorage.removeItem('userData');
    currentUser = null;
    authToken = null;
    updateAuthUI();
    navigateTo('home');
}

// Navigation
function navigateTo(page, data = {}) {
    const links = document.querySelectorAll('.sidebar-nav a');
    links.forEach(link => {
        link.classList.remove('active');
        if (link.dataset.page === page) {
            link.classList.add('active');
        }
    });

    const content = document.getElementById('mainContent');
    content.innerHTML = '<div class="loading">Memuat...</div>';

    switch(page) {
        case 'home':
            loadHomePage();
            break;
        case 'bookmark':
            loadbookmarkPage();
            break;
        case 'following':
            loadFollowingPage();
            break;
        case 'create':
            loadCreateArticlePage();
            break;
        case 'article':
            loadArticleDetailPage(data.id);
            break;
        case 'author':
            loadAuthorPage(data.id);
            break;
        case 'profile':
            if (currentUser) {
                loadAuthorPage(currentUser.id);
            } else {
                navigateTo('login');
            }
            break;
        case 'login':
            loadLoginPage();
            break;
        case 'register':
            loadRegisterPage();
            break;
        default:
            loadHomePage();
    }
}

// Home Page
async function loadHomePage() {
    try {
        const response = await fetch(`${API_BASE}/article/all`);
        const article = await response.json();
        
        const content = document.getElementById('mainContent');
        content.innerHTML = `
            <h1 class="page-title">Artikel Terbaru</h1>
            <div class="article-list" id="articleList"></div>
        `;
        
        renderArticleList(article);
    } catch (error) {
        showError('Gagal memuat artikel');
    }
}

function renderArticleList(article) {
    const listEl = document.getElementById('articleList');
    
    if (!article || article.length === 0) {
        listEl.innerHTML = `
            <div class="empty-state">
                <h3>Belum ada artikel</h3>
                <p>Mulai menulis artikel pertama Anda!</p>
            </div>
        `;
        return;
    }
    
    listEl.innerHTML = article.map(article => `
        <article class="article-card">
            <div class="article-author">
                <div class="author-avatar">${getInitials(article.author_name || 'User')}</div>
                <div class="author-info">
                    <a href="#" class="author-name" onclick="navigateTo('author', {id: ${article.author_id}}); return false;">
                        ${article.author_name || 'Anonymous'}
                    </a>
                    <div class="article-date">${formatDate(article.created_at)}</div>
                </div>
            </div>
            <div class="article-content" onclick="navigateTo('article', {id: ${article.id}})">
                <h2 class="article-title">${article.title}</h2>
                <p class="article-excerpt">${article.content ? article.content.substring(0, 150) + '...' : ''}</p>
            </div>
            <div class="article-meta">
                <span>${article.read_time || 5} menit baca</span>
                <div class="article-actions">
                    <button class="icon-btn ${article.is_bookmarked ? 'active' : ''}" onclick="toggleBookmark(${article.id}, this)">
                        <svg width="20" height="20" viewBox="0 0 24 24" fill="${article.is_bookmarked ? 'currentColor' : 'none'}" stroke="currentColor" stroke-width="2">
                            <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/>
                        </svg>
                    </button>
                </div>
            </div>
        </article>
    `).join('');
}

// Article Detail Page
async function loadArticleDetailPage(articleId) {
    try {
        const response = await fetch(`${API_BASE}/article/${articleId}`);
        const article = await response.json();
        
        const content = document.getElementById('mainContent');
        content.innerHTML = `
            <article>
                <div class="article-detail-header">
                    <h1 class="article-detail-title">${article.title}</h1>
                    <div class="article-author">
                        <div class="author-avatar">${getInitials(article.author_name || 'User')}</div>
                        <div class="author-info">
                            <a href="#" class="author-name" onclick="navigateTo('author', {id: ${article.author_id}}); return false;">
                                ${article.author_name || 'Anonymous'}
                            </a>
                            <div class="article-date">${formatDate(article.created_at)} · ${article.read_time || 5} menit baca</div>
                        </div>
                    </div>
                </div>
                <div class="article-detail-content">
                    ${formatContent(article.content)}
                </div>
            </article>
        `;
    } catch (error) {
        showError('Gagal memuat artikel');
    }
}

// Author Page
async function loadAuthorPage(authorId) {
    try {
        const [userResponse, articleResponse] = await Promise.all([
            fetch(`${API_BASE}/user/${authorId}`),
            fetch(`${API_BASE}/article?author_id=${authorId}`)
        ]);
        
        const user = await userResponse.json();
        const article = await articleResponse.json();
        
        const isOwnProfile = currentUser && currentUser.id === authorId;
        const isFollowing = user.is_following || false;
        
        const content = document.getElementById('mainContent');
        content.innerHTML = `
            <div class="profile-header">
                <div class="profile-info">
                    <div class="profile-avatar">${getInitials(user.name || user.username)}</div>
                    <div class="profile-details">
                        <h1 class="profile-name">${user.name || user.username}</h1>
                        <p class="profile-bio">${user.bio || 'Belum ada bio'}</p>
                        <div class="profile-stats">
                            <span>${article.length} Artikel</span>
                            <span>${user.followers_count || 0} Followers</span>
                            <span>${user.following_count || 0} Following</span>
                        </div>
                        ${!isOwnProfile ? `
                            <button class="btn ${isFollowing ? 'btn-secondary' : 'btn-primary'}" 
                                    id="followBtn" 
                                    onclick="toggleFollow(${authorId}, this)">
                                ${isFollowing ? 'Unfollow' : 'Follow'}
                            </button>
                        ` : ''}
                    </div>
                </div>
            </div>
            <h2 class="page-title">Artikel</h2>
            <div class="article-list" id="articleList"></div>
        `;
        
        renderArticleList(article);
    } catch (error) {
        showError('Gagal memuat profil');
    }
}

// bookmark Page
async function loadbookmarkPage() {
    if (!currentUser) {
        navigateTo('login');
        return;
    }
    
    try {
        const response = await fetch(`${API_BASE}/bookmark/all`, {
            headers: {
                'Authorization': `Bearer ${authToken}`
            }
        });
        const bookmark = await response.json();
        
        const content = document.getElementById('mainContent');
        content.innerHTML = `
            <h1 class="page-title">Bookmark</h1>
            <div class="article-list" id="articleList"></div>
        `;
        
        renderArticleList(bookmark);
    } catch (error) {
        showError('Gagal memuat bookmark');
    }
}

// Following Page
async function loadFollowingPage() {
    if (!currentUser) {
        navigateTo('login');
        return;
    }
    
    try {
        const response = await fetch(`${API_BASE}/follow/all`, {
            headers: {
                'Authorization': `Bearer ${authToken}`
            }
        });
        const article = await response.json();
        
        const content = document.getElementById('mainContent');
        content.innerHTML = `
            <h1 class="page-title">Following</h1>
            <div class="article-list" id="articleList"></div>
        `;
        
        renderArticleList(article);
    } catch (error) {
        showError('Gagal memuat artikel');
    }
}

// Create Article Page
function loadCreateArticlePage() {
    if (!currentUser) {
        navigateTo('login');
        return;
    }
    
    const content = document.getElementById('mainContent');
    content.innerHTML = `
        <div class="form-container">
            <h1 class="page-title">Tulis Artikel Baru</h1>
            <form id="articleForm" onsubmit="handleCreateArticle(event)">
                <div class="form-group">
                    <label class="form-label">Judul</label>
                    <input type="text" class="form-input" name="title" required>
                </div>
                <div class="form-group">
                    <label class="form-label">Konten</label>
                    <textarea class="form-textarea" name="content" required></textarea>
                </div>
                <div class="form-group">
                    <label class="form-label">Topik</label>
                    <select class="form-select" name="topic_id" id="topicSelect">
                        <option value="">Pilih topik...</option>
                    </select>
                </div>
                <button type="submit" class="btn btn-primary">Publikasikan</button>
            </form>
        </div>
    `;
    
    loadTopics();
}

async function loadTopics() {
    try {
        const response = await fetch(`${API_BASE}/topic/all`);
        const topics = await response.json();
        
        const select = document.getElementById('topicSelect');
        select.innerHTML = '<option value="">Pilih topik...</option>' +
            topics.map(topic => `<option value="${topic.id}">${topic.name}</option>`).join('');
    } catch (error) {
        console.error('Gagal memuat topik');
    }
}

async function handleCreateArticle(event) {
    event.preventDefault();
    
    const formData = new FormData(event.target);
    const data = {
        title: formData.get('title'),
        content: formData.get('content'),
        topic_id: parseInt(formData.get('topic_id'))
    };
    
    try {
        const response = await fetch(`${API_BASE}/article/create`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${authToken}`
            },
            body: JSON.stringify(data)
        });
        
        if (response.ok) {
            const article = await response.json();
            navigateTo('article', {id: article.id});
        } else {
            showError('Gagal membuat artikel');
        }
    } catch (error) {
        showError('Gagal membuat artikel');
    }
}

// Login Page
function loadLoginPage() {
    const content = document.getElementById('mainContent');
    content.innerHTML = `
        <div class="auth-container">
            <h1 class="auth-title">Masuk ke Medium</h1>
            <form id="loginForm" onsubmit="handleLogin(event)">
                <div class="form-group">
                    <label class="form-label">Email atau Username</label>
                    <input type="text" class="form-input" name="username" required>
                </div>
                <div class="form-group">
                    <label class="form-label">Password</label>
                    <input type="password" class="form-input" name="password" required>
                </div>
                <button type="submit" class="btn btn-primary" style="width: 100%;">Masuk</button>
            </form>
            <div class="auth-switch">
                Belum punya akun? <a href="#" onclick="navigateTo('register'); return false;">Daftar</a>
            </div>
        </div>
    `;
}

async function handleLogin(event) {
    event.preventDefault();
    
    const formData = new FormData(event.target);
    const data = {
        username: formData.get('username'),
        password: formData.get('password')
    };
    
    try {
        const response = await fetch(`${API_BASE}/user/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        
        if (response.ok) {
            const result = await response.json();
            authToken = result.token;
            currentUser = result.user;
            
            localStorage.setItem('authToken', authToken);
            localStorage.setItem('userData', JSON.stringify(currentUser));
            
            updateAuthUI();
            navigateTo('home');
        } else {
            showError('Login gagal. Periksa username dan password Anda.');
        }
    } catch (error) {
        showError('Terjadi kesalahan saat login');
    }
}

// Register Page
function loadRegisterPage() {
    const content = document.getElementById('mainContent');
    content.innerHTML = `
        <div class="auth-container">
            <h1 class="auth-title">Daftar ke Medium</h1>
            <form id="registerForm" onsubmit="handleRegister(event)">
                <div class="form-group">
                    <label class="form-label">Username</label>
                    <input type="text" class="form-input" name="username" required>
                </div>
                <div class="form-group">
                    <label class="form-label">Email</label>
                    <input type="email" class="form-input" name="email" required>
                </div>
                <div class="form-group">
                    <label class="form-label">Password</label>
                    <input type="password" class="form-input" name="password" required>
                </div>
                <button type="submit" class="btn btn-primary" style="width: 100%;">Daftar</button>
            </form>
            <div class="auth-switch">
                Sudah punya akun? <a href="#" onclick="navigateTo('login'); return false;">Masuk</a>
            </div>
        </div>
    `;
}

async function handleRegister(event) {
    event.preventDefault();
    
    const formData = new FormData(event.target);
    const data = {
        username: formData.get('username'),
        email: formData.get('email'),
        password: formData.get('password')
    };
    
    try {
        const response = await fetch(`${API_BASE}/user/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });
        
        if (response.ok) {
            alert('Registrasi berhasil! Silakan login.');
            navigateTo('login');
        } else {
            showError('Registrasi gagal. Username atau email mungkin sudah digunakan.');
        }
    } catch (error) {
        showError('Terjadi kesalahan saat registrasi');
    }
}

// Toggle Functions
async function toggleBookmark(articleId, btnElement) {
    if (!currentUser) {
        navigateTo('login');
        return;
    }
    
    try {
        const method = btnElement.classList.contains('active') ? 'DELETE' : 'POST';
        const response = await fetch(`${API_BASE}/bookmark/${articleId}`, {
            method: method,
            headers: {
                'Authorization': `Bearer ${authToken}`
            }
        });
        
        if (response.ok) {
            btnElement.classList.toggle('active');
            const svg = btnElement.querySelector('svg');
            svg.setAttribute('fill', btnElement.classList.contains('active') ? 'currentColor' : 'none');
        }
    } catch (error) {
        console.error('Gagal toggle bookmark');
    }
}

async function toggleFollow(userId, btnElement) {
    if (!currentUser) {
        navigateTo('login');
        return;
    }
    
    try {
        const isFollowing = btnElement.textContent.trim() === 'Unfollow';
        const method = isFollowing ? 'DELETE' : 'POST';
        
        const response = await fetch(`${API_BASE}/following/${userId}`, {
            method: method,
            headers: {
                'Authorization': `Bearer ${authToken}`
            }
        });
        
        if (response.ok) {
            btnElement.textContent = isFollowing ? 'Follow' : 'Unfollow';
            btnElement.className = isFollowing ? 'btn btn-primary' : 'btn btn-secondary';
        }
    } catch (error) {
        console.error('Gagal toggle follow');
    }
}

// Utility Functions
function getInitials(name) {
    return name.split(' ').map(n => n[0]).join('').substring(0, 2).toUpperCase();
}

function formatDate(dateString) {
    if (!dateString) return '';
    const date = new Date(dateString);
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    return date.toLocaleDateString('id-ID', options);
}

function formatContent(content) {
    if (!content) return '';
    return content.split('\n').map(p => `<p>${p}</p>`).join('');
}

function showError(message) {
    const content = document.getElementById('mainContent');
    content.innerHTML = `
        <div class="empty-state">
            <h3>Terjadi Kesalahan</h3>
            <p>${message}</p>
            <button class="btn btn-primary" onclick="navigateTo('home')">Kembali ke Beranda</button>
        </div>
    `;
}
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: #fff;
            line-height: 1.6;
            padding: 20px;
        }

        .container {
            max-width: 1200px;
            margin: 0 auto;
            background: rgba(255, 255, 255, 0.1);
            backdrop-filter: blur(10px);
            border-radius: 30px;
            padding: 60px 40px;
            box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
        }

        .header {
            text-align: center;
            margin-bottom: 60px;
        }

        .logo {
            width: 150px;
            height: 150px;
            margin: 0 auto 20px;
            animation: float 3s ease-in-out infinite;
        }

        @keyframes float {
            0%, 100% { transform: translateY(0px); }
            50% { transform: translateY(-20px); }
        }

        h1 {
            font-size: 4em;
            margin: 20px 0;
            background: linear-gradient(45deg, #00d4ff, #00ff88);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            background-clip: text;
            text-shadow: 0 0 30px rgba(0, 212, 255, 0.5);
        }

        .tagline {
            font-size: 1.5em;
            margin-bottom: 30px;
            opacity: 0.9;
        }

        .badges {
            display: flex;
            justify-content: center;
            gap: 15px;
            flex-wrap: wrap;
            margin: 30px 0;
        }

        .badge {
            background: rgba(255, 255, 255, 0.2);
            padding: 8px 20px;
            border-radius: 25px;
            font-size: 0.9em;
            border: 2px solid rgba(255, 255, 255, 0.3);
            transition: all 0.3s;
        }

        .badge:hover {
            background: rgba(255, 255, 255, 0.3);
            transform: translateY(-3px);
            box-shadow: 0 10px 20px rgba(0, 0, 0, 0.2);
        }

        .features {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
            gap: 30px;
            margin: 50px 0;
        }

        .feature-card {
            background: rgba(255, 255, 255, 0.15);
            padding: 30px;
            border-radius: 20px;
            text-align: center;
            border: 2px solid rgba(255, 255, 255, 0.2);
            transition: all 0.3s;
        }

        .feature-card:hover {
            transform: translateY(-10px);
            background: rgba(255, 255, 255, 0.25);
            box-shadow: 0 15px 40px rgba(0, 0, 0, 0.3);
        }

        .feature-icon {
            font-size: 3em;
            margin-bottom: 15px;
        }

        .feature-card h3 {
            font-size: 1.5em;
            margin-bottom: 10px;
        }

        .code-block {
            background: rgba(0, 0, 0, 0.4);
            padding: 30px;
            border-radius: 15px;
            margin: 40px 0;
            border-left: 5px solid #00ff88;
            font-family: 'Courier New', monospace;
            overflow-x: auto;
        }

        .code-block code {
            color: #00ff88;
            font-size: 1em;
        }

        .projects {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
            gap: 25px;
            margin: 50px 0;
        }

        .project-card {
            background: linear-gradient(135deg, rgba(0, 212, 255, 0.2), rgba(0, 255, 136, 0.2));
            padding: 30px;
            border-radius: 20px;
            border: 2px solid rgba(255, 255, 255, 0.3);
            transition: all 0.3s;
            position: relative;
            overflow: hidden;
        }

        .project-card::before {
            content: '';
            position: absolute;
            top: -50%;
            left: -50%;
            width: 200%;
            height: 200%;
            background: linear-gradient(45deg, transparent, rgba(255, 255, 255, 0.1), transparent);
            transform: rotate(45deg);
            transition: all 0.6s;
        }

        .project-card:hover::before {
            left: 100%;
        }

        .project-card:hover {
            transform: scale(1.05);
            box-shadow: 0 20px 50px rgba(0, 0, 0, 0.4);
        }

        .project-emoji {
            font-size: 2.5em;
            margin-bottom: 15px;
        }

        .project-card h3 {
            font-size: 1.4em;
            margin-bottom: 10px;
        }

        .project-tags {
            font-size: 0.85em;
            opacity: 0.8;
            margin-top: 15px;
        }

        .cta-section {
            text-align: center;
            margin: 60px 0;
            padding: 40px;
            background: rgba(0, 0, 0, 0.3);
            border-radius: 20px;
        }

        .cta-button {
            display: inline-block;
            padding: 15px 40px;
            background: linear-gradient(45deg, #00d4ff, #00ff88);
            color: #000;
            text-decoration: none;
            border-radius: 30px;
            font-weight: bold;
            font-size: 1.2em;
            margin: 10px;
            transition: all 0.3s;
            border: none;
            cursor: pointer;
        }

        .cta-button:hover {
            transform: scale(1.1);
            box-shadow: 0 10px 30px rgba(0, 212, 255, 0.5);
        }

        .footer {
            text-align: center;
            margin-top: 60px;
            padding-top: 30px;
            border-top: 2px solid rgba(255, 255, 255, 0.2);
        }

        .gif-container {
            text-align: center;
            margin: 30px 0;
        }

        .gif-container img {
            border-radius: 15px;
            max-width: 400px;
            width: 100%;
            box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
        }

        @media (max-width: 768px) {
            h1 { font-size: 2.5em; }
            .container { padding: 40px 20px; }
            .features { grid-template-columns: 1fr; }
            .projects { grid-template-columns: 1fr; }
        }
    </style>

<body>
    <div class="container">
        <div class="header">
            <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original-wordmark.svg" alt="Go Logo" class="logo">
            <h1>🔥 Damn-Go</h1>
            <p class="tagline">Master Golang from Zero to Hero</p>
            
            <div class="badges">
                <span class="badge">🚀 Go 1.21+</span>
                <span class="badge">📚 20+ Modules</span>
                <span class="badge">💼 4 Real Projects</span>
                <span class="badge">⭐ MIT Licensed</span>
            </div>
        </div>

        <div class="gif-container">
            <img src="https://media.giphy.com/media/qgQUggAC3Pfv687qPC/giphy.gif" alt="Coding Animation">
        </div>

        <div class="features">
            <div class="feature-card">
                <div class="feature-icon">🎓</div>
                <h3>Learn Core Go</h3>
                <p>From basics to concurrency, master all fundamentals</p>
            </div>
            <div class="feature-card">
                <div class="feature-icon">🌐</div>
                <h3>Build Web Apps</h3>
                <p>REST APIs, WebSockets & real-time applications</p>
            </div>
            <div class="feature-card">
                <div class="feature-icon">🚀</div>
                <h3>Deploy Anywhere</h3>
                <p>Docker, databases & production best practices</p>
            </div>
        </div>

        <div class="cta-section">
            <h2>⚡ Get Started in 60 Seconds</h2>
            <div class="code-block">
                <code>

# Clone the repository<br>

git clone https://github.com/madhav9757/damn-go.git<br>
cd damn-go/01-basics<br>
<br>

# Run your first Go program<br>

go run hello-world.go
</code>
</div>
</div>

        <h2 style="text-align: center; font-size: 2.5em; margin: 50px 0 30px;">🛠️ Real-World Projects</h2>

        <div class="projects">
            <div class="project-card">
                <div class="project-emoji">🛒</div>
                <h3>E-commerce API</h3>
                <p>Full backend with auth, payments & order management</p>
                <div class="project-tags">PostgreSQL • JWT • Stripe</div>
            </div>

            <div class="project-card">
                <div class="project-emoji">🤖</div>
                <h3>Discord Bot</h3>
                <p>Automate servers with custom commands & events</p>
                <div class="project-tags">Discord API • Goroutines</div>
            </div>

            <div class="project-card">
                <div class="project-emoji">📊</div>
                <h3>GitHub Analyzer</h3>
                <p>Repository insights & developer statistics</p>
                <div class="project-tags">GitHub API • Data Viz</div>
            </div>

            <div class="project-card">
                <div class="project-emoji">💬</div>
                <h3>Chat Server</h3>
                <p>Real-time messaging with WebSocket support</p>
                <div class="project-tags">WebSockets • Redis • Concurrency</div>
            </div>
        </div>

        <div class="cta-section">
            <h2>Ready to Dive In?</h2>
            <p style="margin: 20px 0;">Join hundreds of developers mastering Go</p>
            <a href="https://github.com/madhav9757/damn-go" class="cta-button">⭐ Star on GitHub</a>
            <a href="https://github.com/madhav9757/damn-go/fork" class="cta-button">🍴 Fork & Learn</a>
        </div>

        <div class="gif-container">
            <img src="https://media.giphy.com/media/LmNwrBhejkK9EFP504/giphy.gif" alt="Success Animation" style="max-width: 300px;">
        </div>

        <div class="footer">
            <p style="font-size: 1.2em;">Made with ❤️ by <a href="https://github.com/madhav9757" style="color: #00ff88; text-decoration: none;">madhav9757</a></p>
            <p style="opacity: 0.7; margin-top: 10px;">Perfect for beginners & intermediate developers</p>
        </div>
    </div>

</body>

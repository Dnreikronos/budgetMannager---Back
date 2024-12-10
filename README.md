<h1>Budget Manager</h1>
<p>A backend project to manage users, bills, and budgets, built with <strong>Golang</strong> and JWT-based authentication.</p>

<h2>Features</h2>
<ul>
    <li>JWT-based user authentication</li>
    <li>RESTful APIs for managing users, bills, and budgets</li>
    <li>Database support: PostgreSQL (production) and SQLite (development/testing)</li>
    <li>Secure and scalable API with Gin and GORM</li>
</ul>

<h2>Tech Stack</h2>
<ul>
    <li><strong>Golang</strong> - Core language</li>
    <li><strong>Gin</strong> - Web framework</li>
    <li><strong>GORM</strong> - ORM for database management</li>
    <li><strong>JWT</strong> - User authentication</li>
    <li><strong>PostgreSQL</strong> - Production database</li>
    <li><strong>SQLite</strong> - Development/testing database</li>
</ul>


<h2>Routes</h2>

<h3>User Authentication</h3>
<ul>
    <li><code>POST /register</code> - Register a new user</li>
    <li><code>POST /login</code> - User login and token generation</li>
    <li><code>GET /profile</code> - View user profile (requires JWT)</li>
</ul>

<h3>Bills Management</h3>
<ul>
    <li><code>POST /CreateBill</code> - Create a new bill</li>
    <li><code>PUT /Bill/:id</code> - Update an existing bill</li>
    <li><code>DELETE /Bill/:id</code> - Delete a bill</li>
    <li><code>GET /Bill/:id</code> - Retrieve a specific bill</li>
    <li><code>GET /Bills</code> - List all bills</li>
</ul>

<h3>Budgets Management</h3>
<ul>
    <li><code>POST /CreateBudget</code> - Create a new budget</li>
    <li><code>PUT /Budget/:id</code> - Update an existing budget</li>
    <li><code>DELETE /Budget/:id</code> - Delete a budget</li>
    <li><code>GET /Budget/:id</code> - Retrieve a specific budget</li>
    <li><code>GET /Budgets</code> - List all budgets</li>
</ul>

<h2>Testing</h2>
<p>Run unit tests using:</p>
<pre><code>go test ./...</code></pre>



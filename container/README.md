<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Database Lifecycle Infographic</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700&display=swap" rel="stylesheet">
    <script>
      tailwind.config = {
        theme: {
          extend: {
            colors: {
              'db-bg': '#F4F7F6',
              'db-text': '#212529',
              'db-card': '#FFFFFF',
              'db-primary': '#36A2EB',
              'db-secondary': '#FFC300',
              'db-accent1': '#FF5733',
              'db-accent2': '#4BC0C0',
              'db-danger': '#C70039',
            }
          }
        }
      }
    </script>
    <style>
      body { font-family: 'Inter', sans-serif; }
      .chart-container {
        position: relative;
        width: 100%;
        max-width: 600px;
        margin-left: auto;
        margin-right: auto;
        height: 300px;
        max-height: 400px;
      }
      @media (min-width: 768px) {
        .chart-container { height: 350px; }
      }
      .flow-box {
        border: 2px solid;
        background-color: #FFFFFF;
        padding: 1rem;
        border-radius: 0.5rem;
        text-align: center;
        box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
        width: 90%;
        margin-left: auto;
        margin-right: auto;
        height: 100%;
      }
      .flow-arrow {
        font-size: 2.5rem;
        font-weight: bold;
        margin: 0.5rem 0;
        text-align: center;
        flex-shrink: 0;
      }
    </style>
</head>
<body class="bg-db-bg text-db-text">

    <div class="container mx-auto p-4 md:p-8">

        <header class="text-center mb-12">
            <h1 class="text-4xl md:text-5xl font-bold text-db-primary mb-4">The Go Database Lifecycle</h1>
            <p class="text-xl text-gray-700 max-w-3xl mx-auto">Visualizing how database schemas and data are managed from Development to Production in a Go project.</p>
        </header>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">

            <div class="md:col-span-2 bg-db-card rounded-lg shadow-md p-6">
                <h2 class="text-2xl font-bold text-db-accent1 mb-3">The Core Concept: Schema Migrations</h2>
                <p class="text-gray-700">Database management in modern applications revolves around <strong class="text-db-accent1">schema migrations</strong>. A migration is a version-controlled, auditable script (usually plain SQL) that defines a precise change to the database structure. Instead of manually editing tables, developers run 'up' migrations to apply new changes (like `CREATE TABLE`) and 'down' migrations to revert them (like `DROP TABLE`). This ensures every environment's database structure is consistent and reproducible.</p>
            </div>

            <div class="md:col-span-2 bg-db-card rounded-lg shadow-md p-6">
                <h2 class="text-2xl font-bold text-db-primary mb-4 text-center">The Three-Environment Flow</h2>
                <p class="text-gray-700 text-center mb-8 max-w-2xl mx-auto">The database's role and data change drastically as code moves from a developer's laptop to the live server. The migration scripts are the constant that ties them all together.</p>

                <div class="flex flex-col md:flex-row justify-between items-center md:items-stretch space-y-4 md:space-y-0 md:space-x-4">

                    <div class="flow-box border-db-primary flex-1 flex flex-col">
                        <h3 class="text-2xl font-bold text-db-primary mb-2">1. Development</h3>
                        <p class="text-gray-700 flex-grow">Developers run migrations ('up') on a local database (e.g., a Docker container). They populate it with <strong class="text-db-primary">seed data</strong> (a small set of realistic fake data) to build and test features. This data is ephemeral and can be reset at any time.</p>
                    </div>

                    <div class="flow-arrow text-db-primary">
                        <span class="hidden md:inline">→</span>
                        <span class="md:hidden">↓</span>
                    </div>

                    <div class="flow-box border-db-accent2 flex-1 flex flex-col">
                        <h3 class="text-2xl font-bold text-db-accent2 mb-2">2. CI / Testing</h3>
                        <p class="text-gray-700 flex-grow">The CI pipeline (e.g., GitHub Actions) spins up a <strong class="text-db-accent2">fresh, empty database</strong>. It runs all migrations ('up') to build the schema from scratch. It may run tests with seed data, or use fixtures. After tests pass, the database is completely destroyed. <strong class="text-db-accent2">No data is ever persisted here.</strong></p>
                    </div>

                    <div class="flow-arrow text-db-accent2">
                        <span class="hidden md:inline">→</span>
                        <span class="md:hidden">↓</span>
                    </div>

                    <div class="flow-box border-db-danger flex-1 flex flex-col">
                        <h3 class="text-2xl font-bold text-db-danger mb-2">3. Production</h3>
                        <p class="text-gray-700 flex-grow">This is the live database containing real user data. When deploying, 'up' migrations are run <strong class="text-db-danger">one time</strong> to apply new schema changes. <strong class="text-db-danger">Data is never, ever destroyed.</strong> 'Down' migrations are rarely, if ever, run in production as it could mean data loss. Backups are critical.</p>
                    </div>
                </div>
            </div>

            <div class="bg-db-card rounded-lg shadow-md p-6">
                <h2 class="text-2xl font-bold text-db-accent2 mb-3">Data Persistence: What Stays?</h2>
                <p class="text-gray-700 mb-4">This chart directly answers whether data or just schemas are persisted. "Persisted" means the data lives on after a task. In CI, nothing persists. In production, <strong class="text-db-danger">only real data</strong> persists; seed data is never introduced.</p>
                <div class="chart-container">
                    <canvas id="persistenceChart"></canvas>
                </div>
            </div>

            <div class="bg-db-card rounded-lg shadow-md p-6">
                <h2 class="text-2xl font-bold text-db-secondary mb-3">The Go Migration Toolset</h2>
                <p class="text-gray-700 mb-4">While you *can* use native Go, most projects use dedicated tools to manage migration files and execution. These tools fall into a few common categories.</p>
                <div class="chart-container" style="max-width: 400px; height: 300px;">
                    <canvas id="toolsChart"></canvas>
                </div>
                <h3 class="text-lg font-semibold text-db-text mt-6 mb-2">Common Libraries & Tools:</h3>
                <ul class="list-disc list-inside text-gray-700 space-y-1">
                    <li><strong class="text-db-primary">golang-migrate/migrate</strong>: A very popular, database-agnostic CLI and library.</li>
                    <li><strong class="text-db-accent2">gorm (ORM)</strong>: The Gorm ORM has built-in auto-migration and migration features.</li>
                    <li><strong class="text-db-accent1">sql-migrate</strong>: A database-driven migration tool, popular with Gorm users.</li>
                    <li><strong class="text-db-secondary">pressly/goose</strong>: Another popular CLI and library for managing SQL migrations.</li>
                </ul>
            </div>

            <div class="md:col-span-2 bg-db-card rounded-lg shadow-md p-6">
                <h2 class="text-2xl font-bold text-db-accent1 mb-4">Native Go vs. Migration Libraries</h2>
                <p class="text-gray-700 mb-6">Can this be done with just plain SQL and native Go? <strong class="text-db-accent1">Yes, absolutely.</strong> Libraries just provide convenience. Here’s the trade-off.</p>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div class="border border-db-secondary rounded-lg p-4">
                        <h3 class="text-xl font-bold text-db-secondary mb-3">Approach 1: Native Go + SQL Files</h3>
                        <p class="text-gray-700 mb-4">You manually create `.sql` files (e.g., `001_init.sql`, `002_add_users.sql`) and write a small Go program using the native `database/sql` package to read these files and execute them against the database. You'd also need to create a table (e.g., `schema_versions`) to track which scripts have already been run.</p>
                        <ul class="space-y-2">
                            <li class="flex items-start"><span class="text-green-500 font-bold mr-2">✓</span><span><strong>Pro:</strong> No external dependencies. Full control.</span></li>
                            <li class="flex items-start"><span class="text-green-500 font-bold mr-2">✓</span><span><strong>Pro:</strong> Excellent for learning the fundamentals.</span></li>
                            <li class="flex items-start"><span class="text-red-500 font-bold mr-2">✗</span><span><strong>Con:</strong> You must build the "which script ran" logic yourself.</span></li>
                            <li class="flex items-start"><span class="text-red-500 font-bold mr-2">✗</span><span><strong>Con:</strong> No built-in 'down' migration, locking, or CLI.</span></li>
                        </ul>
                    </div>

                    <div class="border border-db-primary rounded-lg p-4">
                        <h3 class="text-xl font-bold text-db-primary mb-3">Approach 2: Use a Library/Tool</h3>
                        <p class="text-gray-700 mb-4">You use a tool like `golang-migrate/migrate`. You still write the plain `.sql` files (e.g., `001_init.up.sql`, `001_init.down.sql`), but the tool provides the CLI, 'up'/'down' logic, and automatically manages the `schema_versions` table for you.</p>
                        <ul class="space-y-2">
                            <li class="flex items-start"><span class="text-green-500 font-bold mr-2">✓</span><span><strong>Pro:</strong> Handles all the boilerplate logic for you.</span></li>
                            <li class="flex items-start"><span class="text-green-500 font-bold mr-2">✓</span><span><strong>Pro:</strong> Provides a robust CLI (e.g., `migrate -up`, `migrate -down 1`).</span></li>
                            <li class="flex items-start"><span class="text-green-500 font-bold mr-2">✓</span><span><strong>Pro:</strong> Standard, battle-tested, and team-friendly.</span></li>
                            <li class="flex items-start"><span class="text-red-500 font-bold mr-2">✗</span><span><strong>Con:</strong> Adds an external dependency to your project.</span></li>
                        </ul>
                    </div>
                </div>
            </div>

            <div class="md:col-span-2 text-center mt-8">
                <p class="text-lg text-gray-700"><strong>Key Takeaway:</strong> Consistency is the goal. Migration tools ensure the same, version-controlled schema changes are applied everywhere, while the *data* they operate on remains strictly isolated to its proper environment.</p>
            </div>
        </div>

    </div>

    <script>
        window.onload = () => {
            const chartColors = {
                primary: '#36A2EB',
                secondary: '#FFC300',
                accent1: '#FF5733',
                accent2: '#4BC0C0',
                danger: '#C70039',
                text: '#212529'
            };

            const tooltipTitleCallback = (tooltipItems) => {
                const item = tooltipItems[0];
                let label = item.chart.data.labels[item.dataIndex];
                if (Array.isArray(label)) {
                  return label.join(' ');
                } else {
                  return label;
                }
            };

            const persistenceCtx = document.getElementById('persistenceChart');
            if (persistenceCtx) {
                new Chart(persistenceCtx.getContext('2d'), {
                    type: 'bar',
                    data: {
                        labels: ['Development', 'CI / Testing', 'Production'],
                        datasets: [
                            {
                                label: 'Schema',
                                data: [1, 1, 1],
                                backgroundColor: chartColors.primary,
                                borderColor: chartColors.primary,
                                borderWidth: 1
                            },
                            {
                                label: 'Seed/Test Data',
                                data: [1, 1, 0],
                                backgroundColor: chartColors.accent2,
                                borderColor: chartColors.accent2,
                                borderWidth: 1
                            },
                            {
                                label: 'Real (Prod) Data',
                                data: [0, 0, 1],
                                backgroundColor: chartColors.danger,
                                borderColor: chartColors.danger,
                                borderWidth: 1
                            }
                        ]
                    },
                    options: {
                        responsive: true,
                        maintainAspectRatio: false,
                        plugins: {
                            tooltip: {
                                callbacks: {
                                    title: (tooltipItems) => {
                                        return tooltipItems[0].dataset.label;
                                    },
                                    label: (tooltipItem) => {
                                        const value = tooltipItem.raw;
                                        const label = tooltipItem.chart.data.labels[tooltipItem.dataIndex];
                                        return `${label}: ${value === 1 ? 'Yes / Persisted' : 'No / Ephemeral'}`;
                                    }
                                }
                            },
                            legend: {
                                position: 'top',
                            }
                        },
                        scales: {
                            y: {
                                beginAtZero: true,
                                max: 1,
                                ticks: {
                                    stepSize: 1,
                                    callback: function(value) {
                                        return value === 1 ? 'Persisted / Used' : 'Ephemeral / Not Used';
                                    }
                                }
                            }
                        }
                    }
                });
            }

            const toolsCtx = document.getElementById('toolsChart');
            if (toolsCtx) {
                new Chart(toolsCtx.getContext('2d'), {
                    type: 'doughnut',
                    data: {
                        labels: [
                            ['Standalone CLIs', '(e.g., golang-migrate)'],
                            ['ORM-based Migrations', '(e.g., gorm auto-migrate)'],
                            'Native SQL Scripts'
                        ],
                        datasets: [{
                            label: 'Tooling Approach',
                            data: [55, 30, 15],
                            backgroundColor: [
                                chartColors.primary,
                                chartColors.accent2,
                                chartColors.secondary,
                            ],
                            hoverOffset: 4
                        }]
                    },
                    options: {
                        responsive: true,
                        maintainAspectRatio: false,
                        plugins: {
                            tooltip: {
                                callbacks: {
                                    title: tooltipTitleCallback
                                }
                            },
                            legend: {
                                position: 'bottom',
                                labels: {
                                    font: {
                                        size: 10
                                    },
                                    generateLabels: function(chart) {
                                        const labels = Chart.defaults.plugins.legend.labels.generateLabels(chart);
                                        labels.forEach(label => {
                                            if (Array.isArray(label.text)) {
                                                label.text = label.text.join(' ');
                                            }
                                        });
                                        return labels;
                                    }
                                }
                            }
                        }
                    }
                });
            }
        };
    </script>
</body>
</html>

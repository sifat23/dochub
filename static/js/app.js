const knowledgeBase = [
    {
        title: "Git Commit",
        type: "command",
        content: "git commit -m \"Your commit message here\"",
        tags: ["git", "version control"]
    },
    {
        title: "Docker Run",
        type: "command",
        content: "docker run -d -p 8080:80 --name my-container nginx",
        tags: ["docker", "containers"]
    },
    {
        title: "JavaScript Array Map",
        type: "documentation",
        content: "The map() method creates a new array populated with the results of calling a provided function on every element in the calling array.\n\nExample:\nconst numbers = [1, 2, 3];\nconst doubled = numbers.map(num => num * 2);\n// doubled is [2, 4, 6]",
        tags: ["javascript", "arrays"]
    },
    {
        title: "Python Virtual Environment",
        type: "documentation",
        content: "To create a Python virtual environment:\n\npython -m venv myenv\n\nTo activate on Windows:\nmyenv\\Scripts\\activate\n\nTo activate on Unix/MacOS:\nsource myenv/bin/activate",
        tags: ["python", "environment"]
    },
    {
        title: "System Information",
        type: "command",
        content: "uname -a  # Show system information\n\nlsb_release -a  # Show Linux distribution info\n\ndf -h  # Show disk space usage\n\nfree -h  # Show memory usage",
        tags: ["linux", "system"]
    },
    {
        title: "SSH Connection",
        type: "prompt",
        content: "ssh username@hostname -p port_number\n\nExample:\nssh john@example.com -p 2222",
        tags: ["ssh", "networking"]
    }
];

$(document).ready(function() {
    // Search functionality
    function performSearch(query) {
        const resultsContainer = $('#resultsContainer');
        resultsContainer.empty();

        if (!query) {
            resultsContainer.html(`
                        <div class="text-center empty-state py-5">
                            <i class="bi bi-search" style="font-size: 3rem;"></i>
                            <p class="mt-3">Your search results will appear here</p>
                        </div>
                    `);
            return;
        }

        const lowerQuery = query.toLowerCase();
        const results = knowledgeBase.filter(item =>
            item.title.toLowerCase().includes(lowerQuery) ||
            item.content.toLowerCase().includes(lowerQuery) ||
            item.tags.some(tag => tag.toLowerCase().includes(lowerQuery))
        );

        if (results.length === 0) {
            resultsContainer.html(`
                        <div class="text-center empty-state py-5">
                            <i class="bi bi-exclamation-circle" style="font-size: 3rem;"></i>
                            <p class="mt-3">No results found for "${query}"</p>
                        </div>
                    `);
            return;
        }

        results.forEach((result, index) => {
            const delayClass = `delay-${(index % 3) + 1}`;
            resultsContainer.append(`
                        <div class="result-card mb-4 p-4 ${delayClass}">
                            <div class="d-flex justify-content-between align-items-center mb-3">
                                <h5 class="mb-0">${result.title}</h5>
                                <span class="badge bg-dark result-type">${result.type}</span>
                            </div>
                            <div class="result-content">${result.content}</div>
                            <div class="mt-3">
                                ${result.tags.map(tag => `<span class="badge bg-secondary me-1">${tag}</span>`).join('')}
                            </div>
                        </div>
                    `);
        });
    }

    $('#searchButton').click(function() {
        performSearch($('#searchInput').val());
    });

    $('#searchInput').keypress(function(e) {
        if (e.which === 13) {
            performSearch($('#searchInput').val());
        }
    });

    // Initial random results display
    function showRandomResults() {
        const shuffled = [...knowledgeBase].sort(() => 0.5 - Math.random());
        const randomResults = shuffled.slice(0, 3);

        const resultsContainer = $('#resultsContainer');
        resultsContainer.empty();

        randomResults.forEach((result, index) => {
            const delayClass = `delay-${(index % 3) + 1}`;
            resultsContainer.append(`
                        <div class="result-card mb-4 p-4 ${delayClass}">
                            <div class="d-flex justify-content-between align-items-center mb-3">
                                <h5 class="mb-0">${result.title}</h5>
                                <span class="badge bg-dark result-type">${result.type}</span>
                            </div>
                            <div class="result-content">${result.content}</div>
                            <div class="mt-3">
                                ${result.tags.map(tag => `<span class="badge bg-secondary me-1">${tag}</span>`).join('')}
                            </div>
                        </div>
                    `);
        });
    }

    // Show random results on page load
    showRandomResults();
});
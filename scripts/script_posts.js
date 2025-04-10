// Écoute l'envoi du formulaire
document.getElementById("postForm").addEventListener("submit", async (e) => {
    e.preventDefault(); // Empêche le rechargement de la page

    // Récupération des valeurs des champs
    const titleElement = document.getElementById("title");
    const contentElement = document.getElementById("content");
    const title = titleElement.value.trim();
    const content = contentElement.value.trim();

    // Vérifie si les champs ne sont pas vides
    if (!title || !content) {
        alert("Veuillez remplir tous les champs.");
        return;
    }

    try {
        // Envoie la requête POST pour créer un post
        const response = await fetch("/create_post", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                user_id ,          // Valeur en dur à adapter plus tard
                category_id ,      // Idem
                title,
                content
            }),
        });

        const result = await response.text();
        alert(result);

        // Réinitialise le formulaire
        titleElement.value = "";
        contentElement.value = "";

        // Recharge les posts
        loadPosts();

    } catch (error) {
        console.error("Erreur lors de l'envoi du post :", error);
        alert("Une erreur est survenue. Veuillez réessayer.");
    }
});

// Fonction pour charger et afficher les posts
async function loadPosts() {
    const container = document.getElementById("posts-container");
    container.innerHTML = "<p>Chargement des posts...</p>";

    try {
        const res = await fetch("/api/posts");
        const posts = await res.json();

        if (!Array.isArray(posts) || posts.length === 0) {
            container.innerHTML = "<p>Aucun post trouvé.</p>";
            return;
        }

        // Vide le conteneur et ajoute les posts
        container.innerHTML = "";
        posts.forEach(post => {
            const postElement = document.createElement("div");
            postElement.classList.add("post");
            postElement.innerHTML = `
                <h2>${post.title}</h2>
                <p><strong>Auteur :</strong> ${post.username} | <strong>Date :</strong> ${post.createdat}</p>
                <p>${post.content}</p>
                <hr>
            `;
            container.appendChild(postElement);
        });

    } catch (error) {
        console.error("Erreur lors du chargement des posts :", error);
        container.innerHTML = "<p>Erreur lors du chargement des posts.</p>";
    }
}

// Chargement automatique des posts au démarrage
window.addEventListener("DOMContentLoaded", loadPosts);

document.getElementById("postForm").addEventListener("submit", async (e) => { //empêche le rechargement de la page
    e.preventDefault();
    
    //récupère les éléments des champs du formulaire
    const titleElement = document.getElementById("title");  
    const contentElement = document.getElementById("content");

    //récupère les valeurs des champs
    const title = titleElement.value;
    const content = contentElement.value;

    //envoie une requête POST à l'API pour créer un post
    const response = await fetch("/create_post", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ user_id: 1, category_id: 1, title, content }),
    });

    const result = await response.text();
    alert(result);

    titleElement.value = ""; //vide le champ titre
    contentElement.value = ""; //vide le champ contenu
});
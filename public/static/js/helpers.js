function addHashtags() {
    // Get the current text from the textarea
    let text = document.getElementById("hashtags").value;

    // Split the text by spaces or commas
    let tags = text.split(/[ ,]+/);

    // Ensure each tag starts with a '#'
    let validTags = tags.map(tag => {
      return tag.startsWith('#') ? tag : `#${tag.trim()}`;
    });

    // Update the textarea value with the modified hashtags
    document.getElementById("hashtags").value = validTags.join(' ');
    
  }

  // Attach the event listener to the textarea for real-time updates
  document.getElementById("hashtags").addEventListener("input", addHashtags);
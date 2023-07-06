package main

import "snippetbox.reganloper.com/internal/models"

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the build progresses.
// Also include a Snippets field in the templateData struct to
// hold a slice of snippets (i.e. for home page).
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

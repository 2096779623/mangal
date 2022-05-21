package tui

import (
	"fmt"
	"github.com/metafates/mangai/shared"
)

// plural transforms singular into plural if needed
func plural(count int, word string) string {
	if count == 1 {
		return word
	}

	return word + "s"
}

// View renders the UI.
func (b Bubble) View() string {
	var rendered string

	switch b.state {
	case searchState:
		rendered = renderSearchState(b)
	case spinnerState:
		rendered = renderSpinnerState(b)
	case mangaSelectState:
		rendered = renderMangaSelectState(b)
	case chaptersSelectState:
		rendered = renderChaptersSelectState(b)
	case promptState:
		rendered = renderPromptState(b)
	case progressState:
		rendered = renderProgressState(b)
	case exitPromptState:
		rendered = renderExitPromptState(b)
	}

	return commonStyle.Render(rendered)
}

func renderSearchState(b Bubble) string {
	return fmt.Sprintf(`%s

%s

%s
`, titleStyle.Render(shared.Mangai), b.input.View(), b.help.View(b.stateKeyMap()))
}

func renderSpinnerState(b Bubble) string {
	return fmt.Sprintf(`%s %s

%s`, b.spinner.View(), "Searching...", b.help.View(b.stateKeyMap()))
}

func renderMangaSelectState(b Bubble) string {
	return b.manga.View()
}

func renderChaptersSelectState(b Bubble) string {
	return b.chapters.View()
}

func renderPromptState(b Bubble) string {
	count := len(b.selected)

	return fmt.Sprintf(`Download %d %s of %s?

%s`, count, plural(count, "chapter"), b.prevManga, b.help.View(b.stateKeyMap()))
}

func renderProgressState(b Bubble) string {
	return fmt.Sprintf(`Downloading %s

%s

%s


%s

%s

%s`,
		b.prevManga,
		b.progress.View(),
		textSecondaryStyle.Render(b.prevChapter),
		b.subProgress.View(),
		textSecondaryStyle.Render(b.prevPanel),
		b.help.View(b.stateKeyMap()))
}

func renderExitPromptState(b Bubble) string {
	count := len(b.selected)

	return fmt.Sprintf(`%d %s of %s downloaded

%s`, count, plural(count, "chapter"), b.prevManga, b.help.View(b.stateKeyMap()))
}

package invites

import "time"

// TODO: tags
//	standard tag: WELCOME_CONTENT (must be there)
//  standard tag: RECOMMENDED (must be there)
//  other possible tag: EXERCISE, DIET

// TODO: this assumes a simple assets service that we can post files to & link from
// TODO: ensure links (e.g images) in content that we stage here are cacheable
//	this might involve proxying and rewriting them
// TODO: downsize large images etc
// TODO: consider: content to show during onboarding

type Content struct {
	ID           string
	Title        string
	Body         string // TODO: standardize format e.g HTML or MD
	Author       string
	AuthorAvatar string // TODO: Link; ensure this avatar is cacheable (e.g ETags)
	HeroImage    string // TODO: optional; e.g videos may not have one; also make cacheable
	ContentType  string // TODO: enum e.g video, audio, article

	CreatedAt time.Time
	UpdatedAt time.Time

	// e.g estimated time to read an article or video/audio length
	Estimate int // TODO: standardize unit e.g seconds

	// tags are used to filter and target content
	// some tags are standard (prescribed )
	Tags []Tag
}

// TODO: update content from CMS / fetch from CMC
// TODO: react e.g love | metrics
// TODO: save | metrics
// TODO: share | metrics
// TODO: fetch (sorting e.g by new), limiting (e.g top 3), filtering (e.g tag) | defaults
// TODO: query by more than one tag at a time (union)
// TODO: related content e.g based on tags , limit #
// TODO: each time a single article is fetched, count as a view for metrics
// TODO: welcome content -> a predefined tag

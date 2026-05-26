// Copyright 2023 Trillian Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package claimant is a code model for the Claimant Model.
package claimant

import (
	_ "embed"
)

var (
	// TemplateModelMarkdown is the markdown template for a model.
	//go:embed tmpl_model.md
	TemplateModelMarkdown []byte

	// TemplateQuestionsMarkdown is the markdown template for a model questionnaire.
	//go:embed tmpl_questions.md
	TemplateQuestionsMarkdown []byte

	// TemplateSequenceMarkdown is the markdown template for the sequence diagram.
	//go:embed tmpl_sequence.md
	TemplateSequenceMarkdown []byte
)

// Claim represents a falsifiable statement along with an actor that can verify it.
type Claim struct {
	// Claim is a falsifiable statement, e.g. "The number 7 is prime"
	Claim string `yaml:"Claim"`
	// Verifier is an actor that can verify the claim, e.g. "Primal Calculator"
	Verifier string `yaml:"Verifier"`
}

// Model represents a Claimant Model, mapping the roles to actors.
// See https://github.com/google/trillian/blob/master/docs/claimantmodel/CoreModel.md for
// descriptions of these roles. Repeating the definitions here will only lead to stale
// documentation.
type Model struct {
	// System is a short upper case name that models the essence of this model, e.g. "FIRMWARE".
	System string `yaml:"System"`
	// Claimant is the actor playing the role of the Claimant.
	Claimant string `yaml:"Claimant"`
	// Statement is the concrete type that the Claimant issues, and is likely the thing that is logged.
	Statement string `yaml:"Statement"`

	// Believer is the actor playing the role of the Believer.
	Believer string `yaml:"Believer,omitempty"`
	// Believers are the actor playing the roles of the Believer.
	// This should only be provided if there are multiple Believers, and if
	// provided then Believer should be left empty.
	Believers []string `yaml:"Believers,omitempty"`

	// Claim is the claim made by the Claimant.
	Claim Claim `yaml:"Claim,omitempty"`
	// Claims are the claims made by the Claimant.
	// This should only be provided if there are multiple Claims, and if
	// provided then Claim should be left empty.
	Claims []Claim `yaml:"Claims,omitempty"`

	// Arbiter is the actor or process that fulfills the Arbiter role.
	Arbiter string `yaml:"Arbiter"`
}

// Markdown returns this Claimant Model in a definition table that renders
// clearly in markdown format.
func (m Model) Markdown() string { _ = "STUB: not implemented"; return "" }

// Questionnaire returns some questions to guide the designer to ensure that
// the claimant model is sound.
func (m Model) Questionnaire() string { _ = "STUB: not implemented"; return "" }

// ClaimTerms finds all of the terms used in the Claim that must be
// present in the Statement.
func (m Model) ClaimTerms() []string { _ = "STUB: not implemented"; return nil }

// ClaimMarkdown renders the Claim(s) in markdown.
func (m Model) ClaimMarkdown() string { _ = "STUB: not implemented"; return "" }

// VerifierList returns all of the verifiers mapped to the claim they verify.
func (m Model) VerifierList() map[string]string { _ = "STUB: not implemented"; return nil }

// VerifierMarkdown renders the Verifier(s) in markdown.
func (m Model) VerifierMarkdown() string { _ = "STUB: not implemented"; return "" }

// BelieverMarkdown renders the Believer(s) in markdown.
func (m Model) BelieverMarkdown() string { _ = "STUB: not implemented"; return "" }

// LogModelForDomain proposes a template Claimant Model for human
// editing based on a domain model provided.
func LogModelForDomain(m Model) Model { _ = "STUB: not implemented"; return *new(Model) }

// Models captures the domain model along with the log model that supports it.
// This can be extended for more general model composition in the future, but
// this is the most common composition and motiviation for Claimant Modelling.
type Models struct {
	Domain Model `yaml:"Domain"`
	Log    Model `yaml:"Log"`
}

// Actors returns all of the actors that participate in the ecosystem of logging
// the domain claims and verifying all behaviours.
func (ms Models) Actors() []string { _ = "STUB: not implemented"; return nil }

// TODO(mhutchinson): put these in a more useful order than alphabetical

// Markdown returns the markdown representation of both models.
func (ms Models) Markdown() string { _ = "STUB: not implemented"; return "" }

// SequenceDiagram returns a mermaid markdown snippet that shows the
// idealized workflow for this log ecosystem. This can be changed in the
// future to support other variations in the workflow, e.g. this generates
// a sequence that shows the claimant awaiting an inclusion proof and then
// creating an offline bundle, but not all ecosystems do this and so perhaps
// this should take some kind of Options that allows these cases to vary.
// For now, this is out of scope and this generated sequence diagram should
// be taken to represent the current best practice, and designers can modify
// it to reflect the deltas in their world.
func (ms Models) SequenceDiagram() string { _ = "STUB: not implemented"; return "" }

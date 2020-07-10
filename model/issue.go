// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

type Issue struct {
	RepoOwner string
	RepoName  string
	Number    int
	Username  string
	State     string
	Labels    []string
}

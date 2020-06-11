package common

import (
	"github.com/zerefwayne/college-portal-backend/course"
	"github.com/zerefwayne/college-portal-backend/resource"
)

// Common consists of a usecase instances which can be used throughout the project without
// causing the problem of cyclic imports
var (
	Course   course.Usecase
	Resource resource.Usecase
)

package models

// NavItem model for navigation bar
type NavItem struct {
	Path     string
	Title    string
	Selected bool
}

// NavigationItems structure
type NavigationItems struct {
	NavItems []NavItem
}

// Selected sets the selected state of the navigation items
func (navItems *NavigationItems) Selected(path string) {
	for idx := range navItems.NavItems {
		navItems.NavItems[idx].Selected = false
		if navItems.NavItems[idx].Path == path {
			navItems.NavItems[idx].Selected = true
		}
	}
}

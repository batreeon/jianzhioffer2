func replaceSpaces(str string) string {
    r := strings.NewReplacer(" ", "%20")
    return r.Replace(str)
}
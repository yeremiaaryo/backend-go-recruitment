package movie

func (ma *movieApplication) InsertLogs(keyword, pagination string) error {
	return ma.movieRepository.Insert(keyword, pagination)
}

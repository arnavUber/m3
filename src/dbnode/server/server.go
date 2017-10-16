	"path"
	"github.com/m3db/m3/src/x/lockfile"
	filePathPrefixLockFile           = ".lock"
	err := cfg.InitDefaultsAndValidate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initializing config defaults and validating config: %v", err)
		os.Exit(1)
	}

	newFileMode, err := cfg.Filesystem.ParseNewFileMode()
	if err != nil {
		logger.Fatalf("could not parse new file mode: %v", err)
	}

	newDirectoryMode, err := cfg.Filesystem.ParseNewDirectoryMode()
	if err != nil {
		logger.Fatalf("could not parse new directory mode: %v", err)
	}

	// Obtain a lock on `filePathPrefix`, or exit if another process already has it.
	// The lock consists of a lock file (on the file system) and a lock in memory.
	// When the process exits gracefully, both the lock file and the lock will be removed.
	// If the process exits ungracefully, only the lock in memory will be removed, the lock
	// file will remain on the file system. When a dbnode starts after an ungracefully stop,
	// it will be able to acquire the lock despite the fact the the lock file exists.
	lockPath := path.Join(cfg.Filesystem.FilePathPrefixOrDefault(), filePathPrefixLockFile)
	fslock, err := lockfile.CreateAndAcquire(lockPath, newDirectoryMode)
	if err != nil {
		logger.Fatalf("could not acquire lock on %s: %v", lockPath, err)
	}
	defer fslock.Release()

			SetLimitMbps(cfg.Filesystem.ThroughputLimitMbpsOrDefault()).
			SetLimitCheckEvery(cfg.Filesystem.ThroughputCheckEveryOrDefault())).
	// Setup postings list cache.
	var (
		plCacheConfig  = cfg.Cache.PostingsListConfiguration()
		plCacheSize    = plCacheConfig.SizeOrDefault()
		plCacheOptions = index.PostingsListCacheOptions{
			InstrumentOptions: opts.InstrumentOptions().
				SetMetricsScope(scope.SubScope("postings-list-cache")),
		}
	)
	postingsListCache, stopReporting, err := index.NewPostingsListCache(plCacheSize, plCacheOptions)
	if err != nil {
		logger.Fatalf("could not construct query cache: %s", err.Error())
	}
	defer stopReporting()

	indexOpts = indexOpts.SetInsertMode(insertMode).
		SetPostingsListCache(postingsListCache).
		SetReadThroughSegmentOptions(index.ReadThroughSegmentOptions{
			CacheRegexp: plCacheConfig.CacheRegexpOrDefault(),
			CacheTerms:  plCacheConfig.CacheTermsOrDefault(),
		})
	opts = opts.SetIndexOptions(indexOpts)
	mmapCfg := cfg.Filesystem.MmapConfigurationOrDefault()
		poolOptions(
			policy.TagEncoderPool,
			scope.SubScope("tag-encoder-pool")))
		poolOptions(
			policy.TagDecoderPool,
			scope.SubScope("tag-decoder-pool")))
		SetFilePathPrefix(cfg.Filesystem.FilePathPrefixOrDefault()).
		SetWriterBufferSize(cfg.Filesystem.WriteBufferSizeOrDefault()).
		SetDataReaderBufferSize(cfg.Filesystem.DataReadBufferSizeOrDefault()).
		SetInfoReaderBufferSize(cfg.Filesystem.InfoReadBufferSizeOrDefault()).
		SetSeekReaderBufferSize(cfg.Filesystem.SeekReadBufferSizeOrDefault()).
		SetForceIndexSummariesMmapMemory(cfg.Filesystem.ForceIndexSummariesMmapMemoryOrDefault()).
		SetForceBloomFilterMmapMemory(cfg.Filesystem.ForceBloomFilterMmapMemoryOrDefault())
		// Feature currently not working.
		SetRepairEnabled(false)
		b.Capacity = bucket.CapacityOrDefault()
		b.Count = bucket.SizeOrDefault()
			SetRefillLowWatermark(bucket.RefillLowWaterMarkOrDefault()).
			SetRefillHighWatermark(bucket.RefillHighWaterMarkOrDefault())
	switch policy.TypeOrDefault() {
		poolOptions(
			policy.SegmentReaderPool,
			scope.SubScope("segment-reader-pool")))

		poolOptions(
			policy.EncoderPool,
			scope.SubScope("encoder-pool")))

	closersPoolOpts := poolOptions(
		policy.ClosersPool,
		scope.SubScope("closers-pool"))

	contextPoolOpts := poolOptions(
		policy.ContextPool.PoolPolicy,
		scope.SubScope("context-pool"))

		SetMaxPooledFinalizerCapacity(policy.ContextPool.MaxFinalizerCapacityOrDefault()))

		poolOptions(
			policy.IteratorPool,
			scope.SubScope("iterator-pool")))

		poolOptions(
			policy.IteratorPool,
			scope.SubScope("multi-iterator-pool")))
	var writeBatchPoolSize int
	if policy.WriteBatchPool.Size != nil {
		writeBatchPoolSize = *policy.WriteBatchPool.Size
	} else {
		// If no value set, calculate a reasonable value based on the commit log
		writeBatchPoolSize = commitlogQueueSize / expectedBatchSize

	writeBatchPoolOpts := pool.NewObjectPoolOptions()
	writeBatchPoolOpts = writeBatchPoolOpts.
		SetSize(writeBatchPoolSize).
		// Set watermarks to zero because this pool is sized to be as large as we
		// ever need it to be, so background allocations are usually wasteful.
		SetRefillLowWatermark(0.0).
		SetRefillHighWatermark(0.0).
		SetInstrumentOptions(
			writeBatchPoolOpts.
				InstrumentOptions().
				SetMetricsScope(scope.SubScope("write-batch-pool")))

		writeBatchPoolOpts,
	tagPoolPolicy := policy.TagsPool
		IDPoolOptions: poolOptions(
			policy.IdentifierPool, scope.SubScope("identifier-pool")),
		TagsPoolOptions: maxCapacityPoolOptions(tagPoolPolicy, scope.SubScope("tags-pool")),
		TagsCapacity:    tagPoolPolicy.CapacityOrDefault(),
		TagsMaxCapacity: tagPoolPolicy.MaxCapacityOrDefault(),
		TagsIteratorPoolOptions: poolOptions(
			policy.TagsIteratorPool,
			scope.SubScope("tags-iterator-pool")),
	fetchBlockMetadataResultsPoolPolicy := policy.FetchBlockMetadataResultsPool
		capacityPoolOptions(
			fetchBlockMetadataResultsPoolPolicy,
		fetchBlockMetadataResultsPoolPolicy.CapacityOrDefault())
	fetchBlocksMetadataResultsPoolPolicy := policy.FetchBlocksMetadataResultsPool
		capacityPoolOptions(
			fetchBlocksMetadataResultsPoolPolicy,
		fetchBlocksMetadataResultsPoolPolicy.CapacityOrDefault())
		SetDatabaseBlockAllocSize(policy.BlockAllocSizeOrDefault()).
	blockPool := block.NewDatabaseBlockPool(
		poolOptions(
			policy.BlockPool,
			scope.SubScope("block-pool")))
		poolOptions(
			policy.SeriesPool,
			scope.SubScope("series-pool")))
	var (
		resultsPool = index.NewResultsPool(
			poolOptions(policy.IndexResultsPool, scope.SubScope("index-results-pool")))
		postingsListOpts = poolOptions(policy.PostingsListPool, scope.SubScope("postingslist-pool"))
		postingsList     = postings.NewPool(postingsListOpts, roaring.NewPostingsList)
	)


	var (
		opts                = pool.NewObjectPoolOptions()
		size                = policy.SizeOrDefault()
		refillLowWaterMark  = policy.RefillLowWaterMarkOrDefault()
		refillHighWaterMark = policy.RefillHighWaterMarkOrDefault()
	)

	if size > 0 {
		opts = opts.SetSize(size)
		if refillLowWaterMark > 0 &&
			refillHighWaterMark > 0 &&
			refillHighWaterMark > refillLowWaterMark {
			opts = opts.
				SetRefillLowWatermark(refillLowWaterMark).
				SetRefillHighWatermark(refillHighWaterMark)
	var (
		opts                = pool.NewObjectPoolOptions()
		size                = policy.SizeOrDefault()
		refillLowWaterMark  = policy.RefillLowWaterMarkOrDefault()
		refillHighWaterMark = policy.RefillHighWaterMarkOrDefault()
	)

	if size > 0 {
		opts = opts.SetSize(size)
		if refillLowWaterMark > 0 &&
			refillHighWaterMark > 0 &&
			refillHighWaterMark > refillLowWaterMark {
			opts = opts.SetRefillLowWatermark(refillLowWaterMark)
			opts = opts.SetRefillHighWatermark(refillHighWaterMark)
	var (
		opts                = pool.NewObjectPoolOptions()
		size                = policy.SizeOrDefault()
		refillLowWaterMark  = policy.RefillLowWaterMarkOrDefault()
		refillHighWaterMark = policy.RefillHighWaterMarkOrDefault()
	)

	if size > 0 {
		opts = opts.SetSize(size)
		if refillLowWaterMark > 0 &&
			refillHighWaterMark > 0 &&
			refillHighWaterMark > refillLowWaterMark {
			opts = opts.SetRefillLowWatermark(refillLowWaterMark)
			opts = opts.SetRefillHighWatermark(refillHighWaterMark)
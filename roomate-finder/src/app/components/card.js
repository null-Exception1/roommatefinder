export default function Card({ BlockID, PartialCount, FullCount }) {
  return (
    <a
      href={`/block/${BlockID}`}
      className="card w-full sm:w-72 border-2 border-black bg-white shadow-md hover:shadow-xl transition-transform hover:scale-105"
    >
      <div className="card-body text-center space-y-3">
        <h2 className="card-title text-lg sm:text-xl">BLOCK {BlockID}</h2>
        <div className="flex flex-wrap justify-center gap-2">
          <span className="inline-block text-xs sm:text-sm bg-red-500 text-white px-2 py-1 rounded">
            {FullCount} FULL
          </span>
          <span className="inline-block text-xs sm:text-sm bg-orange-500 text-white px-2 py-1 rounded">
            {PartialCount} PARTIAL
          </span>
        </div>
      </div>
    </a>
  );
}

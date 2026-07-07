
export default function Card({ BlockID, PartialCount, FullCount }) {
  return (
    <a href={`/block/${BlockID}`} className="card hover:bg-blue-500 w-80 border-2 border-black bg-white card-xl shadow-md hover:shadow-2xl">
      <div className="card-body text-center">
        <h2 className="card-title">BLOCK {BlockID}</h2>
        <div className="flex flex-wrap justify-center gap-2">
          <span className="inline-block text-sm bg-red-500 text-white px-2 py-1 rounded">
            {FullCount} FULL
          </span>
          <span className="inline-block text-sm bg-orange-500 text-white px-2 py-1 rounded">

            {PartialCount} PARTIAL
          </span>
        </div>
      </div>
    </a>
  );
}
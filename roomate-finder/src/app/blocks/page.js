"use client"
import Navbar from "@/components/navbar";
import Card from "@/components/card";
import Breadcumbs from "@/components/breadcrumbs";
import { useEffect, useState } from "react";

export default function Blocks() {
  const baseURL = process.env.NEXT_PUBLIC_API_URL;
  const url = `${baseURL}/blocks`;

  const [loaded, setLoaded] = useState(false);
  const [loaded_data, setLoadedData] = useState([]);

  useEffect(() => {
    fetch(url)
      .then(response => response.json())
      .then(data => {
        setLoaded(true);
        setLoadedData(data);
        console.log("Data:", loaded, data);
      });
  }, []);

  return (
    <div className="flex justify-center">
      <Navbar />

      <div className="flex flex-col m-4 sm:m-20 flex-wrap w-full sm:w-11/12 border-2 border-black rounded-3xl">
        <div className="pt-6 sm:pt-10 pl-4 sm:pl-10">
          <fieldset className="fieldset max-w-xs text-black">
            <legend className="fieldset-legend text-black">Campus</legend>
            <select
              defaultValue="Main Campus"
              className="select bg-white text-black border-gray-300 focus:border-blue-500"
            >
              <option disabled={true}>Pick a campus</option>
              <option className="bg-white text-black">Main Campus</option>
              <option className="bg-white text-black">Coming soon</option>
            </select>
          </fieldset>
        </div>

        <Breadcumbs crumbs={["Home", "Blocks"]} links={["/", "/blocks"]} />

        <h1 className="pl-4 sm:pl-10 pt-2 text-xl sm:text-2xl">Blocks</h1>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 p-4 sm:p-5">
          {loaded ? (
            Object.entries(loaded_data).map(([blockID, value]) => (
              <Card
                key={blockID}
                BlockID={blockID}
                PartialCount={value.Partial}
                FullCount={value.Full}
              />
            ))
          ) : (
            <div className="flex justify-center items-center p-10">
              <span className="loading loading-spinner loading-xl"></span>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

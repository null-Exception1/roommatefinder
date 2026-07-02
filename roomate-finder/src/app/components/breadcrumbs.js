export default function Breadcrumbs({ crumbs, links }) {
  return (
    <div className="breadcrumbs text-sm ml-10 pb-5 pt-10 overflow-hidden">
      <ul className="overflow-hidden">
        {crumbs.map((crumb, index) => (
          <li key={index}>
            {index < crumbs.length - 1 ? (<a href={links[index]} className="text-black-500 hover:underline">{crumb}</a>) : (<span>{crumb}</span>)}
          </li>
        ))}
      </ul>
    </div>
  );
}
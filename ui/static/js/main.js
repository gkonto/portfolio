document.documentElement.classList.add("js");

const navLinks = Array.from(document.querySelectorAll(".nav-links a"));
const sectionLinks = navLinks.filter((link) => link.hash);
const homeLink = navLinks.find((link) => link.getAttribute("href") === "/");
const pathLinks = navLinks.filter((link) => {
  const href = link.getAttribute("href") || "";
  return href.startsWith("/") && !link.hash;
});

const setActiveLink = (targetId = "") => {
  navLinks.forEach((link) => {
    link.classList.remove("is-active");
    link.removeAttribute("aria-current");
  });

  if (window.location.pathname !== "/") {
    const activePathLink = pathLinks.find((link) => link.getAttribute("href") === window.location.pathname);
    if (activePathLink) {
      activePathLink.classList.add("is-active");
      activePathLink.setAttribute("aria-current", "page");
    }
    return;
  }

  if (!targetId && homeLink) {
    homeLink.classList.add("is-active");
    homeLink.setAttribute("aria-current", "page");
    return;
  }

  const activeLink = sectionLinks.find((link) => link.hash === `#${targetId}`);
  if (activeLink) {
    activeLink.classList.add("is-active");
    activeLink.setAttribute("aria-current", "page");
  }
};

const sections = sectionLinks
  .map((link) => document.querySelector(link.hash))
  .filter(Boolean);

if (window.location.pathname === "/") {
  if ("IntersectionObserver" in window && sections.length > 0) {
    const observer = new IntersectionObserver(
      (entries) => {
        const visible = entries
          .filter((entry) => entry.isIntersecting)
          .sort((a, b) => b.intersectionRatio - a.intersectionRatio)[0];

        if (!visible) {
          return;
        }

        setActiveLink(visible.target.id);
      },
      {
        rootMargin: "-20% 0px -55% 0px",
        threshold: [0.2, 0.4, 0.6],
      },
    );

    sections.forEach((section) => observer.observe(section));
  }

  const hashTarget = window.location.hash.replace("#", "");
  setActiveLink(hashTarget);
  window.addEventListener("hashchange", () => {
    setActiveLink(window.location.hash.replace("#", ""));
  });
}

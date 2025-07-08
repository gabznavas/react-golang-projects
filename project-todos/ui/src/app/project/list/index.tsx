'use client'

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { PlusIcon, SearchIcon } from "lucide-react";
import { ProjectCard } from "@/components/project-card";
import { useEffect, useState } from "react";
import { requests } from "@/services/api";
import { Project } from "@/services/types";

export default function ProjectList() {
  const [projects, setProjects] = useState<Project[]>([]);

  useEffect(() => {
    const fetchProjects = async () => {
      const projects = await requests.getProjects();
      setProjects(prev => {
        const newProjects = projects.filter((project: Project) => !prev.some(p => p.id === project.id));
        return [...prev, ...newProjects];
      });
    }
    fetchProjects();
  }, []);

  return (
    <div className="flex flex-col gap-4 mx-auto">
      <div className="flex gap-4 justify-between items-center">
        <Input type='search' placeholder="Procure por um projeto" />
        <Button autoFocus variant="outline" className="flex items-center gap-2 w-1/6">
          <SearchIcon />
        </Button>
        <Button variant="default" className="w-1/10">
          <PlusIcon />
        </Button>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 2xl:grid-cols-4 gap-4 w-full">
        {
          projects.map(project => (
            <ProjectCard key={project.id} project={project} />
          ))
        }
      </div>
    </div>

  )
}
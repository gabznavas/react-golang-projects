import { Header } from "@/components/header";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { requests } from "@/services/api";
import { useRouter, useSearchParams } from "next/navigation";
import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";

interface ProjectFormData {
  name: string;
  description: string;
}

const schema = yup.object().shape({
  name: yup.string().required('Nome é obrigatório'),
  description: yup.string().required('Descrição é obrigatória'),
});

export default function ProjectDetails() {
  const params = useSearchParams();
  const id = params.get('id');
  const router = useRouter();

  const form = useForm<ProjectFormData>({
    resolver: yupResolver(schema),
    defaultValues: {
      name: '',
      description: '',
    }
  });

  useEffect(() => {
    const fetchProject = async () => {
      if (id) {
        const project = await requests.getProjectById(id);
        form.reset({
          name: project.name,
          description: project.description,
        });
      }
    }
    fetchProject();
  }, [id, form]);

  const handleSave = async (formData: ProjectFormData) => {
    const payload = {
      name: formData.name,
      description: formData.description
    }
    if (id) {
      await requests.updateProject(id, payload);
    } else {
      await requests.createProject(payload);
    }
    router.push('/project/list');
  }

  return (
    <div className="flex flex-col gap-4 mx-auto">
      <Header />
      <div className="flex flex-col gap-4 w-[550px] mx-auto mt-10">
        <span className="text-xl font-bold">
          {id ? `Projeto ${form.getValues().name}` : 'Novo Projeto'}
        </span>
        <Form {...form}>
          <form className="flex flex-col gap-4" onSubmit={form.handleSubmit(handleSave)}>
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Nome do projeto</FormLabel>
                  <FormControl>
                    <Input
                      type="text"
                      placeholder="Nome do projeto"
                      {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="description"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Descrição do projeto</FormLabel>
                  <FormControl>
                    <Textarea
                      className="h-[150px]"
                      placeholder="Descrição do projeto"
                      {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <div className="flex gap-4 justify-end">
              <Button variant="default" className="w-2/10 cursor-pointer">
                {id ? 'Atualizar' : 'Criar'}
              </Button>
              <Button variant="outline"
                className="w-2/10 cursor-pointer"
                type="button" onClick={() => router.push('/project/list')} >
                Cancelar
              </Button>
            </div>
          </form>
        </Form>
      </div>
    </div>
  )
}
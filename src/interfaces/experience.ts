import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

class IExperience {

	@prop({ required: true })
	schema_version: number;

	@prop()
	location?: IAddress;

	@prop({ required: true })
	id: string;

	@prop()
	organization?: string;

	@prop()
	opportunity?: string;

	@prop()
	description?: string;

	@prop()
	when?: [string, string];

	@prop({ required: true })
	is_verified: boolean;

	@prop({ required: true })
	created_at: number;
}
